import {
  getGetStockQueryKey,
  getListStocksQueryKey,
  useDeleteStock,
  useGetStock,
  useListStocks,
  useListStockTags,
  useUpdateStock,
} from "@/api/api";
import { Overlay, OverlayLoading } from "@/components/Overlay";
import PageTitle from "@/components/PageTitle";
import { AddBtn } from "@/components/Preference";
import type { Stock } from "@/types/stock";
import { useEffect, useState } from "react";
import type { Tag } from "@/types/tag";
import { CircleX, Trash2 } from "lucide-react";
import { useQueryClient } from "@tanstack/react-query";

export default function PreferenceStocks() {
  const [stockId, setStockId] = useState("");
  const [isShowEditer, setIsShowEditer] = useState(false);
  const [isShowAdder, setIsShowAdder] = useState(false);
  return (
    <>
      {isShowEditer && (
        <StockEditor
          stockId={stockId}
          onClose={() => {
            setIsShowEditer(false);
          }}
        />
      )}
      <div className="flex place-content-between">
        <PageTitle title="Preference-Stocks"></PageTitle>
        <AddBtn
          className="mb-5"
          onAdd={() => {
            setIsShowAdder(true);
          }}
        />
      </div>
      <StocksList
        onStockSelect={(id) => {
          setStockId(id);
          setIsShowEditer(true);
        }}
      />
    </>
  );
}
function StocksList({
  onStockSelect,
}: {
  onStockSelect: (id: string) => void;
}) {
  const { data, isPending, isError } = useListStocks();
  if (isPending) {
    return <OverlayLoading />;
  }
  if (isError || data?.status != 200) {
    alert("Failed to fetch transactions data.");
    return;
  }
  return (
    <div className="flex gap-2 flex-wrap">
      {data.data.map((stock) => (
        <div
          onClick={() => onStockSelect(stock.id)}
          className="border w-xs border-primary-light rounded-xl p-2 flex flex-col gap-x-2 text-xl cursor-pointer hover:bg-primary-lighter"
          key={stock.id}
        >
          <p>{stock.name}</p>
        </div>
      ))}
    </div>
  );
}

function StockEditor({
  stockId,
  onClose,
}: {
  stockId: string;
  onClose: () => void;
}) {
  const { data, isPending, isError } = useGetStock(stockId);
  const {
    data: allTagsData,
    isPending: allTagsIsPending,
    isError: allTagsIsError,
  } = useListStockTags();
  const changeStockMutation = useUpdateStock();
  const deleteStockMutation = useDeleteStock();
  const queryClient = useQueryClient();
  const [stock, setStock] = useState<Stock>({
    name: "",
    hasAmount: true,
    currency: "",
    currencyExponent: 0,
    description: "",
    tags: [],
  });
  useEffect(() => {
    if (data?.status === 200) {
      setStock({
        name: data.data.name,
        hasAmount: data.data.hasAmount,
        currency: data.data.currency,
        currencyExponent: data.data.currencyExponent,
        description: data.data.description,
        tags: data.data.tags,
      });
    }
  }, [data, stockId]);

  if (isPending || allTagsIsPending) {
    return <OverlayLoading />;
  }
  if (
    isError ||
    data.status != 200 ||
    allTagsIsError ||
    allTagsData.status != 200
  ) {
    alert("Failed to fetch transactions data.");
    return;
  }

  function changeStock() {
    changeStockMutation.mutate(
      {
        id: stockId,
        data: stock,
      },
      {
        onSuccess: async () => {
          await queryClient.invalidateQueries({
            queryKey: getGetStockQueryKey(stockId),
          });
          await queryClient.invalidateQueries({
            queryKey: getListStocksQueryKey(),
          });
          onClose();
        },
        onError: () => {
          alert("Failed to change the stock.");
        },
      },
    );
  }
  function deleteStock() {
    deleteStockMutation.mutate(
      {
        id: stockId,
      },
      {
        onSuccess: async () => {
          await queryClient.invalidateQueries({
            queryKey: getListStocksQueryKey(),
          });
          onClose();
        },
        onError: () => {
          alert("Failed to delete the stock.");
        },
      },
    );
  }
  return (
    <Overlay click={onClose}>
      <form
        className="m-auto rounded-xl bg-primary-lighter p-4 flex flex-col w-md"
        onSubmit={(e) => {
          e.preventDefault();
          changeStock();
        }}
      >
        <StockFormParts
          stock={stock}
          setStock={(stock) => setStock(stock)}
          allTags={allTagsData.data}
        />
        <div className="mt-3 flex">
          <button
            type="submit"
            className="cursor-pointer rounded-md p-1 border-2 border-primary-light hover:bg-primary-light flex-1 mr-3 font-bold"
          >
            Change
          </button>
          <button
            onClick={deleteStock}
            type="button"
            className="cursor-pointer rounded-md p-1 border-2 border-primary-light hover:bg-primary-light"
          >
            <Trash2 />
          </button>
        </div>
      </form>
    </Overlay>
  );
}

function StockFormParts({
  stock,
  setStock,
  allTags,
}: {
  stock: Stock;
  setStock: (stock: Stock) => void;
  allTags: Tag[];
}) {
  return (
    <>
      <input
        type="text"
        placeholder="Stock Name"
        value={stock.name}
        onChange={(e) => setStock({ ...stock, name: e.target.value })}
        required
        className="bg-primary-lightest p-2 text-xl rounded-md mb-3"
      />
      <div className="flex mb-3 w-full place-content-between">
        <div className="w-1/4">
          <label htmlFor="hasAmount">Countable</label>
          <input
            id="hasAmount"
            type="checkbox"
            checked={stock.hasAmount}
            onChange={(e) =>
              setStock({ ...stock, hasAmount: e.target.checked })
            }
            required
            className="block m-auto size-8"
          />
        </div>
        <div className="w-1/3">
          <label htmlFor="currency">Currency</label>
          <input
            id="currency"
            type="text"
            placeholder="Currency"
            value={stock.currency}
            onChange={(e) => setStock({ ...stock, currency: e.target.value })}
            required
            className="bg-primary-lightest p-2 rounded-md w-full text-center block"
          />
        </div>
        <div className="w-1/3">
          <label htmlFor="currencyExponent">Currency Exponent</label>
          <input
            id="currencyExponent"
            type="number"
            placeholder="Currency Exponent"
            value={stock.currencyExponent}
            onChange={(e) =>
              setStock({ ...stock, currencyExponent: e.target.valueAsNumber })
            }
            required
            className="bg-primary-lightest p-2 rounded-md w-full text-center block"
          />
        </div>
      </div>
      <textarea
        placeholder="Description"
        value={stock.description}
        onChange={(e) => setStock({ ...stock, description: e.target.value })}
        required
        className="bg-primary-lightest p-2 rounded-md mb-3"
      />
      <TagsEditor
        allTags={allTags}
        tags={stock.tags}
        setTags={(tags) => setStock({ ...stock, tags: tags })}
      />
    </>
  );
}
function TagsEditor({
  allTags,
  tags,
  setTags,
}: {
  allTags: Tag[];
  tags: string[];
  setTags: (tags: string[]) => void;
}) {
  return (
    <div className="flex gap-2">
      {tags.map((id) => (
        <div key={id} className="bg-primary-lightest p-1 rounded-xl">
          <CircleX
            className="mr-0.5 inline align-sub"
            size={18}
            onClick={() => setTags(tags.filter((t) => t != id))}
          />
          {allTags.filter((t) => t.id === id)[0].name}
        </div>
      ))}
      <select
        onChange={(e) => {
          setTags([...tags, e.target.value]);
        }}
        className="bg-primary-lightest p-1 rounded-xl"
      >
        <option value="default">Add Tag</option>
        {allTags
          .filter((t) => !tags.includes(t.id))
          .map((u) => (
            <option value={u.id} key={u.id}>
              {u.name}
            </option>
          ))}
      </select>
    </div>
  );
}
