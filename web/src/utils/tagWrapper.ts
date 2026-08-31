import {
  useListStockTags,
  useListTransactionTags,
  useListFlowTags,
  useUpdateFlowTag,
  useUpdateStockTag,
  useUpdateTransactionTag,
  getListStockTagsQueryKey,
  getListTransactionTagsQueryKey,
  getListFlowTagsQueryKey,
  useDeleteStockTag,
  useDeleteTransactionTag,
  useDeleteFlowTag,
  useCreateStockTag,
  useCreateTransactionTag,
  useCreateFlowTag,
} from "@/api/api";
import type { TagType } from "@/types/tag";

export function useListTags(tagType: TagType) {
  const stock = useListStockTags({
    query: { enabled: tagType == "stock" },
  });
  const transaction = useListTransactionTags({
    query: { enabled: tagType == "transaction" },
  });
  const flow = useListFlowTags({
    query: { enabled: tagType == "flow" },
  });
  switch (tagType) {
    case "stock":
      return stock;
    case "transaction":
      return transaction;
    case "flow":
      return flow;
  }
}

export function useUpdateTag(tagType: TagType) {
  const stock = useUpdateStockTag();
  const transaction = useUpdateTransactionTag();
  const flow = useUpdateFlowTag();
  switch (tagType) {
    case "stock":
      return stock;
    case "transaction":
      return transaction;
    case "flow":
      return flow;
  }
}

export function useDeleteTag(tagType: TagType) {
  const stock = useDeleteStockTag();
  const transaction = useDeleteTransactionTag();
  const flow = useDeleteFlowTag();
  switch (tagType) {
    case "stock":
      return stock;
    case "transaction":
      return transaction;
    case "flow":
      return flow;
  }
}

export function getListTagsQueryKey(tagType: TagType) {
  switch (tagType) {
    case "stock":
      return getListStockTagsQueryKey();
    case "transaction":
      return getListTransactionTagsQueryKey();
    case "flow":
      return getListFlowTagsQueryKey();
  }
}

export function useCreateTag(tagType: TagType) {
  const stock = useCreateStockTag();
  const transaction = useCreateTransactionTag();
  const flow = useCreateFlowTag();
  switch (tagType) {
    case "stock":
      return stock;
    case "transaction":
      return transaction;
    case "flow":
      return flow;
  }
}
