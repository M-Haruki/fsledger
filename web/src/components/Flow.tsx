import type { Flow } from "@/types/flow";
import type { StockAbstract } from "@/types/stock";
import type { Tag } from "@/types/tag";
import { TagsEditor } from "./Tag";
import type { UUID } from "@/types/uuid";
import { ArrowBigRight, Plus, X } from "lucide-react";
import { bigint2string, string2bigint } from "@/utils/bigint";
import { useEffect, useState } from "react";
import { Switch } from "./Switch";

export function Flows({
  flows,
  setFlows,
  allStocks,
  allFlowTags,
}: {
  flows: Flow[];
  setFlows: (flows: Flow[]) => void;
  allStocks: StockAbstract[];
  allFlowTags: Tag[];
}) {
  return (
    <div>
      {flows.map((flow, index) => (
        <Flow
          key={index}
          flow={flow}
          setFlow={(flow) =>
            setFlows(flows.map((f, j) => (index === j ? flow : f)))
          }
          deleteFlow={() => setFlows(flows.filter((_, j) => index !== j))}
          allStocks={allStocks}
          allFlowTags={allFlowTags}
        />
      ))}
      <Plus
        onClick={() =>
          setFlows(
            flows.concat({
              from: "",
              to: "",
              fromAmount: BigInt(0),
              toAmount: BigInt(0),
              tags: [],
            }),
          )
        }
      />
    </div>
  );
}

function Flow({
  flow,
  setFlow,
  deleteFlow,
  allStocks,
  allFlowTags,
}: {
  flow: Flow;
  setFlow: (flow: Flow) => void;
  deleteFlow: () => void;
  allStocks: StockAbstract[];
  allFlowTags: Tag[];
}) {
  const fromCurrencyExponent =
    allStocks.find((stock) => stock.id === flow.from)?.currencyExponent || 0;
  const fromCurrency = allStocks.find(
    (stock) => stock.id === flow.from,
  )?.currency;
  const toCurrencyExponent =
    allStocks.find((stock) => stock.id === flow.to)?.currencyExponent || 0;
  const toCurrency = allStocks.find((stock) => stock.id === flow.to)?.currency;
  const [isSameAmount, setIsSameAmount] = useState(false);
  useEffect(() => {
    setIsSameAmount(fromCurrencyExponent === toCurrencyExponent);
  }, [fromCurrencyExponent, toCurrencyExponent]);

  return (
    <div>
      <X onClick={deleteFlow} />
      <div className="flex">
        <StockSelecter
          stockId={flow.from}
          allStocks={allStocks}
          onChange={(id) => {
            setFlow({ ...flow, from: id });
          }}
        />
        <ArrowBigRight />
        <StockSelecter
          stockId={flow.to}
          allStocks={allStocks}
          onChange={(id) => {
            setFlow({ ...flow, to: id });
          }}
        />
      </div>
      <div className="flex">
        {isSameAmount ? (
          <>
            {fromCurrency == toCurrency
              ? fromCurrency
              : `${fromCurrency}/${toCurrency}`}
            <AmountInput
              amount={flow.fromAmount}
              currencyExponent={fromCurrencyExponent}
              setAmount={(amount) =>
                setFlow({ ...flow, fromAmount: amount, toAmount: amount })
              }
            />
          </>
        ) : (
          <>
            {fromCurrency}
            <AmountInput
              amount={flow.fromAmount}
              currencyExponent={fromCurrencyExponent}
              setAmount={(amount) => setFlow({ ...flow, fromAmount: amount })}
            />
            {toCurrency}
            <AmountInput
              amount={flow.toAmount}
              currencyExponent={toCurrencyExponent}
              setAmount={(amount) => setFlow({ ...flow, toAmount: amount })}
            />
          </>
        )}
      </div>
      <div className="flex">
        <label>Same</label>
        <Switch
          value={isSameAmount}
          onChange={() =>
            fromCurrencyExponent === toCurrencyExponent &&
            setIsSameAmount(!isSameAmount)
          }
        />
      </div>
      <TagsEditor
        tags={flow.tags}
        allTags={allFlowTags}
        setTags={(tags) => {
          setFlow({ ...flow, tags: tags });
        }}
      />
    </div>
  );
}

function StockSelecter({
  stockId,
  allStocks,
  onChange,
}: {
  stockId: UUID;
  allStocks: StockAbstract[];
  onChange: (id: UUID) => void;
}) {
  return (
    <select value={stockId} onChange={(e) => onChange(e.target.value)}>
      {allStocks.map((stock) => (
        <option key={stock.id} value={stock.id}>
          {stock.name}
        </option>
      ))}
      {!allStocks.some((stock) => stock.id === stockId) && (
        <option key={stockId} value={stockId} disabled>
          [Unknown]
        </option>
      )}
    </select>
  );
}

function AmountInput({
  amount,
  currencyExponent,
  setAmount,
}: {
  amount: bigint;
  currencyExponent: number;
  setAmount: (amount: bigint) => void;
}) {
  const [tempAmount, setTempAmount] = useState<string>(
    bigint2string(amount, currencyExponent),
  );
  useEffect(() => {
    setTempAmount(bigint2string(amount, currencyExponent));
  }, [amount, currencyExponent]);
  return (
    <input
      type="text"
      inputMode="decimal"
      value={tempAmount}
      onChange={(e) => {
        const value = e.target.value;
        if (!/^[+-]?\d*\.?\d*$/.test(value)) {
          return;
        }
        setTempAmount(e.target.value);
      }}
      onBlur={(e) => {
        const newAmount = string2bigint(e.target.value, currencyExponent);
        setAmount(newAmount);
        setTempAmount(bigint2string(newAmount, currencyExponent));
      }}
    />
  );
}
