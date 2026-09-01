import type { Flow } from "@/types/flow";
import type { StockAbstract } from "@/types/stock";
import type { Tag } from "@/types/tag";
import { TagsEditor } from "./Tag";

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
  const fromStock = allStocks.find((stock) => flow.from === stock.id);
  const toStock = allStocks.find((stock) => flow.to === stock.id);
  return (
    <div>
      <div>
        <div>{fromStock?.name || "[unlnown]"}</div>
        <div>{flow.amount}</div>
        <div>{toStock?.name || "[unlnown]"}</div>
      </div>
      <TagsEditor
        tags={flow.tags}
        allTags={allFlowTags}
        setTags={(tags) => {}}
      />
    </div>
  );
}
