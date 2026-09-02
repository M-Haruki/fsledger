import type { Flow } from "@/types/flow";
import type { StockAbstract } from "@/types/stock";
import type { Tag } from "@/types/tag";
import { TagsEditor } from "./Tag";
import type { UUID } from "@/types/uuid";
import { ArrowBigRight, Plus, X } from "lucide-react";
import { bigint2string, string2bigint } from "@/utils/bigint";
import { useEffect, useState } from "react";

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
    <div className="flex flex-col gap-y-2">
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
      <div
        className="flex justify-center border-3 rounded-2xl p-1 border-primary-lighter hover:bg-primary-lighter"
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
      >
        <Plus />
      </div>
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
    <div className="border-3 rounded-2xl p-2 border-primary-lighter flex">
      <X onClick={deleteFlow} className="mr-2 cursor-pointer" />
      <div className="flex flex-col w-full gap-y-2">
        <div className="flex w-full gap-x-3">
          <StockSelecter
            className="flex-1 text-lg bg-primary-lightest rounded-md p-1"
            stockId={flow.from}
            allStocks={allStocks}
            onChange={(id) => {
              setFlow({ ...flow, from: id });
            }}
          />
          <ArrowBigRight size={36} />
          <StockSelecter
            className="flex-1 text-lg bg-primary-lightest rounded-md p-1"
            stockId={flow.to}
            allStocks={allStocks}
            onChange={(id) => {
              setFlow({ ...flow, to: id });
            }}
          />
        </div>
        <div className="flex gap-x-5 justify-center">
          {isSameAmount ? (
            <>
              <div className="flex">
                <AmountInput
                  amount={flow.fromAmount}
                  currencyExponent={fromCurrencyExponent}
                  setAmount={(amount) =>
                    setFlow({ ...flow, fromAmount: amount, toAmount: amount })
                  }
                  className="text-lg p-1 rounded-lg bg-primary-lightest w-30 text-right"
                />
                <p className="ml-1 content-center">
                  {fromCurrency == toCurrency
                    ? fromCurrency
                    : `${fromCurrency}/${toCurrency}`}
                </p>
              </div>
            </>
          ) : (
            <>
              <div className="flex">
                <AmountInput
                  amount={flow.fromAmount}
                  currencyExponent={fromCurrencyExponent}
                  setAmount={(amount) =>
                    setFlow({ ...flow, fromAmount: amount })
                  }
                  className="text-lg p-1 rounded-lg bg-primary-lightest w-30 text-right"
                />
                <p className="ml-1 content-center">{fromCurrency}</p>
              </div>
              <div className="flex">
                <AmountInput
                  amount={flow.toAmount}
                  currencyExponent={toCurrencyExponent}
                  setAmount={(amount) => setFlow({ ...flow, toAmount: amount })}
                  className="text-lg p-1 rounded-lg bg-primary-lightest w-30 text-right"
                />
                <p className="ml-1 content-center">{toCurrency}</p>
              </div>
            </>
          )}
        </div>
        <div className="flex">
          <input
            required
            id="isSameAmount"
            type="checkbox"
            checked={isSameAmount}
            onChange={() => setIsSameAmount(!isSameAmount)}
            disabled={fromCurrencyExponent != toCurrencyExponent}
            className="mr-1"
          />
          <label htmlFor="isSameAmount">Use Same Amount</label>
        </div>
        <TagsEditor
          tags={flow.tags}
          allTags={allFlowTags}
          setTags={(tags) => {
            setFlow({ ...flow, tags: tags });
          }}
        />
      </div>
    </div>
  );
}

function StockSelecter({
  stockId,
  allStocks,
  onChange,
  className = "",
}: {
  stockId: UUID;
  allStocks: StockAbstract[];
  onChange: (id: UUID) => void;
  className?: string;
}) {
  return (
    <select
      value={stockId}
      onChange={(e) => onChange(e.target.value)}
      className={className}
    >
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
  className = "",
}: {
  amount: bigint;
  currencyExponent: number;
  setAmount: (amount: bigint) => void;
  className?: string;
}) {
  const [tempAmount, setTempAmount] = useState<string>(
    bigint2string(amount, currencyExponent),
  );
  useEffect(() => {
    setTempAmount(bigint2string(amount, currencyExponent));
  }, [amount, currencyExponent]);
  return (
    <input
      required
      className={className}
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
