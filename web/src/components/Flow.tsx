import type { Flow } from "@/types/flow";
import type { StockAbstract } from "@/types/stock";
import type { Tag } from "@/types/tag";
import { TagsEditor } from "./Tag";
import { Frown } from "lucide-react";

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
      {flows.map((flow, i) => (
        <Flow
          key={i}
          flow={flow}
          setFlow={(flow) => {}}
          allStocks={allStocks}
          allFlowTags={allFlowTags}
        />
      ))}
    </div>
  );
}

function Flow({
  flow,
  setFlow,
  allStocks,
  allFlowTags,
}: {
  flow: Flow;
  setFlow: (flow: Flow) => void;
  allStocks: StockAbstract[];
  allFlowTags: Tag[];
}) {
  return (
    <div>
      <div>
        {/* <div>{fromStock?.name || "[unlnown]"}</div> */}
        <div>{flow.fromAmount}</div>
        <StockSelecter
          stockId={flow.to}
          allStocks={allStocks}
          onChange={(id) => {}}
        />
        -＞
        <div>{flow.toAmount}</div>
        <StockSelecter
          stockId={flow.from}
          allStocks={allStocks}
          onChange={(id) => {}}
        />
      </div>
      <TagsEditor
        tags={flow.tags}
        allTags={allFlowTags}
        setTags={(tags) => {}}
      />
    </div>
  );
}

function StockSelecter({
  stockId,
  allStocks,
  onChange,
}: {
  stockId: string;
  allStocks: StockAbstract[];
  onChange: (id: string) => void;
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
