import type { StockAbstract } from "@/types/stock";
import type { Tag } from "@/types/tag";
import type { Transaction } from "@/types/transaction";
import { Flows } from "./Flow";
import { TagsEditor } from "./Tag";

export function TransactionFormParts({
  transaction,
  setTransaction,
  allStocks,
  allStockTags,
  allFlowTags,
}: {
  transaction: Transaction;
  setTransaction: (transaction: Transaction) => void;
  allStocks: StockAbstract[];
  allStockTags: Tag[];
  allFlowTags: Tag[];
}) {
  return (
    <>
      <div>
        <textarea
          className="bg-primary-lightest p-1 rounded-md w-full"
          value={transaction.description}
          onChange={(e) =>
            setTransaction({ ...transaction, description: e.target.value })
          }
        />
      </div>
      <div>
        <input
          className="bg-primary-lightest p-1 rounded-md text-lg"
          type="date"
          value={transaction.occurredAt}
          onChange={(e) =>
            e.target.value != "" &&
            setTransaction({ ...transaction, occurredAt: e.target.value })
          }
          required
        />
      </div>
      <TagsEditor
        allTags={allStockTags}
        tags={transaction.tags}
        setTags={(tags) => setTransaction({ ...transaction, tags: tags })}
      />
      <Flows
        flows={transaction.flows}
        setFlows={(flows) => {
          setTransaction({ ...transaction, flows: flows });
        }}
        allStocks={allStocks}
        allFlowTags={allFlowTags}
      />
    </>
  );
}
