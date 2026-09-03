import {
  useCreateTransaction,
  useListFlowTags,
  useListStocks,
  useListStockTags,
} from "@/api/api";
import { OverlayLoading } from "@/components/Overlay";
import PageTitle from "@/components/PageTitle";
import { TransactionFormParts } from "@/components/Transaction";
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

  const createTransactionMutation = useCreateTransaction();

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

  function createTransaction() {
    createTransactionMutation.mutate(
      {
        data: {
          description: transaction.description,
          occurred_at: transaction.occurredAt,
          tags: transaction.tags,
          flows: transaction.flows.map((flow) => {
            return {
              from: flow.from,
              to: flow.to,
              from_amount: String(flow.fromAmount),
              to_amount: String(flow.toAmount),
              tags: flow.tags,
            };
          }),
        },
      },
      {
        onSuccess: async () => {},
        onError: (e) => {
          console.log(e);
          alert("Failed to create new transaction.");
        },
      },
    );
  }

  return (
    <>
      <PageTitle title="New Transaction" />
      <form
        className="flex flex-col gap-y-3 w-md"
        onSubmit={(e) => {
          e.preventDefault();
          createTransaction();
        }}
      >
        <TransactionFormParts
          transaction={transaction}
          setTransaction={(transaction) => setTransaction(transaction)}
          allStocks={allStocksData.data}
          allStockTags={allStockTagsData.data}
          allFlowTags={allFlowTagsData.data}
        />
        <button
          type="submit"
          className="cursor-pointer rounded-xl p-1 border-3 border-primary-light hover:bg-primary-light flex-1 font-bold"
        >
          Create
        </button>
      </form>
    </>
  );
}
