import { useListFlowTags, useListStocks, useListStockTags } from "@/api/api";
import { Flows } from "@/components/Flow";
import { OverlayLoading } from "@/components/Overlay";
import PageTitle from "@/components/PageTitle";
import { TagsEditor } from "@/components/Tag";
import type { Transaction } from "@/types/transaction";
import { nowDate } from "@/utils/date";
import { useState } from "react";

export default function TransactionNew() {
  const {
    data: allStocksData,
    isPending: allStocksIsPending,
    isError: allStocksIsError,
  } = useListStocks();
  const {
    data: allFlowTagsData,
    isPending: allFlowTagsIsPending,
    isError: allFlowTagsIsError,
  } = useListFlowTags();
  const {
    data: allStockTagsData,
    isPending: allStockTagsIsPending,
    isError: allStockTagsIsError,
  } = useListStockTags();

  const [transaction, setTransaction] = useState<Transaction>({
    description: "",
    occurredAt: nowDate(),
    tags: [],
    flows: [],
  });

  if (allStocksIsPending || allFlowTagsIsPending || allStockTagsIsPending) {
    return <OverlayLoading />;
  }
  if (
    allStocksIsError ||
    allStocksData.status != 200 ||
    allFlowTagsIsError ||
    allFlowTagsData.status != 200 ||
    allStockTagsIsError ||
    allStockTagsData.status != 200
  ) {
    alert("Failed to fetch required data.");
    return;
  }

  return (
    <>
      <PageTitle title="New Transaction" />
      <div>
        <textarea
          value={transaction.description}
          onChange={(e) =>
            setTransaction({ ...transaction, description: e.target.value })
          }
        />
      </div>
      <div>
        <input
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
        allTags={allStockTagsData.data}
        tags={transaction.tags}
        setTags={(tags) => setTransaction({ ...transaction, tags: tags })}
      />
      <Flows
        flows={transaction.flows}
        setFlows={(flows) => {
          setTransaction({ ...transaction, flows: flows });
        }}
        allStocks={allStocksData.data}
        allFlowTags={allFlowTagsData.data}
      />
    </>
  );
}
