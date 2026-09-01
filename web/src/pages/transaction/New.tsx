import { useListFlowTags, useListStocks } from "@/api/api";
import { Flows } from "@/components/Flow";
import { OverlayLoading } from "@/components/Overlay";
import PageTitle from "@/components/PageTitle";
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

  const [transaction, setTransaction] = useState<Transaction>({
    description: "",
    occurredAt: nowDate(),
    tags: [],
    flows: [],
  });

  if (allStocksIsPending || allFlowTagsIsPending) {
    return <OverlayLoading />;
  }
  if (
    allStocksIsError ||
    allStocksData.status != 200 ||
    allFlowTagsIsError ||
    allFlowTagsData.status != 200
  ) {
    alert("Failed to fetch required data.");
    return;
  }

  return (
    <>
      <PageTitle title="New Transaction" />
      <div>{transaction.description}</div>
      <div>{transaction.occurredAt}</div>
      <Flows
        flows={[
          {
            from: "6bd471ea-37ad-4967-abe6-d84810c5c3bc",
            to: "7debab7d-d3d4-41f0-932c-d7ea8965d7b2",
            fromAmount: BigInt(5000),
            toAmount: BigInt(5000),
            tags: ["a8296388-39e1-4ce1-850c-dcee77d995e1"],
          },
          {
            from: "7debab7d-d3d4-41f0-932c-d7ea8965d7b2",
            to: "6bd471ea-37ad-4967-abe6-d84810c5c3bc",
            fromAmount: BigInt(25),
            toAmount: BigInt(25),
            tags: ["a8296388-39e1-4ce1-850c-dcee77d995e1"],
          },
        ]}
        setFlows={(flows) => {}}
        allStocks={allStocksData.data}
        allFlowTags={allFlowTagsData.data}
      />
    </>
  );
}
