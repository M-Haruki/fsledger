import {
  getGetTransactionQueryKey,
  useDeleteTransaction,
  useGetTransaction,
  useListFlowTags,
  useListStocks,
  useListStockTags,
  useUpdateTransaction,
} from "@/api/api";
import { OverlayLoading } from "@/components/Overlay";
import PageTitle from "@/components/PageTitle";
import { TransactionFormParts } from "@/components/Transaction";
import type { Transaction } from "@/types/transaction";
import { nowDate } from "@/utils/date";
import { useQueryClient } from "@tanstack/react-query";
import { useEffect, useState } from "react";
import { useNavigate, useParams } from "react-router";

export default function TransactionEdit() {
  const transactionId = useParams<{ id: string }>().id;

  const {
    data: transactionData,
    isPending: transactionIsPending,
    isError: transactionIsError,
  } = useGetTransaction(transactionId || "");
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
  const updateTransactionMutation = useUpdateTransaction();
  const deleteTransactionMutation = useDeleteTransaction();
  const queryClient = useQueryClient();
  const navigate = useNavigate();

  const [transaction, setTransaction] = useState<Transaction>({
    description: "",
    occurredAt: nowDate(),
    tags: [],
    flows: [],
  });

  useEffect(() => {
    if (transactionIsError || transactionData?.status != 200) {
      return;
    }
    setTransaction({
      description: transactionData.data.description,
      occurredAt: transactionData.data.occurred_at,
      flows: transactionData.data.flows.map((flow) => {
        return {
          from: flow.from,
          to: flow.to,
          fromAmount: BigInt(flow.from_amount),
          toAmount: BigInt(flow.to_amount),
          tags: flow.tags,
        };
      }),
      tags: transactionData.data.tags,
    });
  }, [transactionIsError, transactionData]);

  if (!transactionId) {
    throw new Error("Missing url param.");
  }

  if (
    transactionIsPending ||
    allStocksIsPending ||
    allFlowTagsIsPending ||
    allStockTagsIsPending
  ) {
    return <OverlayLoading />;
  }
  if (
    transactionIsError ||
    transactionData.status != 200 ||
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

  function updateTransaction() {
    if (!transactionId) {
      throw new Error("Missing url param.");
    }
    updateTransactionMutation.mutate(
      {
        id: transactionId,
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
        onSuccess: async () => {
          await queryClient.invalidateQueries({
            queryKey: getGetTransactionQueryKey(transactionId),
          });
          alert("Changed.");
        },
        onError: () => {
          alert("Failed to update new transaction.");
        },
      },
    );
  }

  function deleteTransaction() {
    if (!transactionId) {
      throw new Error("Missing url param.");
    }
    deleteTransactionMutation.mutate(
      { id: transactionId },
      {
        onSuccess: async () => {
          navigate("/");
        },
        onError: () => {
          alert("Failed to delete new transaction.");
        },
      },
    );
  }

  return (
    <>
      <PageTitle title="Edit Transaction" />
      <form
        className="flex flex-col gap-y-3 w-md"
        onSubmit={(e) => {
          e.preventDefault();
          updateTransaction();
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
          Change
        </button>
        <button
          onClick={deleteTransaction}
          type="button"
          className="cursor-pointer rounded-xl p-1 border-3 border-primary-light hover:bg-primary-light flex-1 font-bold"
        >
          Delete
        </button>
      </form>
    </>
  );
}
