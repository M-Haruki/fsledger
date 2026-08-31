import { useEffect, useState } from "react";
import { Trash2 } from "lucide-react";
import {
  useListTags,
  useUpdateTag,
  getListTagsQueryKey,
  useDeleteTag,
  useCreateTag,
} from "@/utils/tagWrapper.ts";
import PageTitle from "@/components/PageTitle";
import { type TagType, type Tag, NullTag } from "@/types/tag";
import { Overlay, OverlayLoading } from "@/components/Overlay";
import { useQueryClient } from "@tanstack/react-query";
import { AddBtn } from "@/components/Preference";

export default function PreferenceTags({ tagType }: { tagType: TagType }) {
  const [editTag, setEditTag] = useState<Tag>(NullTag);
  const [isShowEditer, setIsShowEditer] = useState(false);
  const [isShowAdder, setIsShowAdder] = useState(false);
  return (
    <>
      {isShowEditer && (
        <TagEditer
          tag={editTag}
          tagType={tagType}
          onClose={() => setIsShowEditer(false)}
        />
      )}
      {isShowAdder && (
        <TagAdder tagType={tagType} onClose={() => setIsShowAdder(false)} />
      )}
      <div className="flex place-content-between">
        <PageTitle title={getPageTitle(tagType)}></PageTitle>
        <AddBtn className="mb-5" onAdd={() => setIsShowAdder(true)} />
      </div>
      <TagList
        tagType={tagType}
        onTagSelect={(tag) => {
          setEditTag(tag);
          setIsShowEditer(true);
        }}
      />
    </>
  );
}

function TagList({
  tagType,
  onTagSelect,
}: {
  tagType: TagType;
  onTagSelect: (tag: Tag) => void;
}) {
  const { data, isPending, isError } = useListTags(tagType);
  if (isPending) {
    return <OverlayLoading />;
  }
  if (isError || data?.status != 200) {
    alert("Failed to fetch tags data.");
    return;
  }
  return (
    <div className="flex gap-2 flex-wrap">
      {data.data.map((tag) => (
        <div
          key={tag.id}
          onClick={() => onTagSelect({ id: tag.id, name: tag.name })}
          className="border w-xs border-primary-light rounded-xl p-2 flex gap-x-2 text-xl cursor-pointer hover:bg-primary-lighter"
        >
          {tag.name}
        </div>
      ))}
    </div>
  );
}

function TagEditer({
  tag,
  tagType,
  onClose,
}: {
  tag: Tag;
  tagType: TagType;
  onClose: () => void;
}) {
  const [tagName, setTagName] = useState("");
  const changeTagMutation = useUpdateTag(tagType);
  const deleteTagMutation = useDeleteTag(tagType);
  const queryClient = useQueryClient();
  useEffect(() => {
    setTagName(tag.name);
  }, [tag]);
  function changeTagName() {
    changeTagMutation.mutate(
      {
        id: tag.id,
        data: { name: tagName },
      },
      {
        onSuccess: async () => {
          await queryClient.invalidateQueries({
            queryKey: getListTagsQueryKey(tagType),
          });
          onClose();
        },
        onError: () => {
          alert("Failed to change the tag name.");
        },
      },
    );
  }
  function deleteTag() {
    deleteTagMutation.mutate(
      { id: tag.id },
      {
        onSuccess: async () => {
          await queryClient.invalidateQueries({
            queryKey: getListTagsQueryKey(tagType),
          });
          onClose();
        },
        onError: () => {
          alert("Failed to delete the tag name.");
        },
      },
    );
  }
  return (
    <Overlay click={onClose}>
      <form
        onSubmit={(e) => {
          e.preventDefault();
          changeTagName();
        }}
        className="m-auto rounded-xl bg-primary-lighter p-4 flex flex-col  w-md"
      >
        <input
          type="text"
          placeholder="Tag Name"
          value={tagName}
          onChange={(e) => setTagName(e.target.value)}
          required
          className="bg-primary-lightest p-2 text-xl rounded-md mb-3"
        />
        <div className="flex">
          <button
            type="submit"
            className="cursor-pointer rounded-md p-1 border-2 border-primary-light hover:bg-primary-light flex-1 mr-3 font-bold"
          >
            Change
          </button>
          <button
            type="button"
            onClick={deleteTag}
            className="cursor-pointer rounded-md p-1 border-2 border-primary-light hover:bg-primary-light"
          >
            <Trash2 />
          </button>
        </div>
      </form>
    </Overlay>
  );
}

function TagAdder({
  tagType,
  onClose,
}: {
  tagType: TagType;
  onClose: () => void;
}) {
  const [tagName, setTagName] = useState("");
  const createTagMutation = useCreateTag(tagType);
  const queryClient = useQueryClient();

  function createTag() {
    createTagMutation.mutate(
      {
        data: {
          name: tagName,
        },
      },
      {
        onSuccess: async () => {
          await queryClient.invalidateQueries({
            queryKey: getListTagsQueryKey(tagType),
          });
          onClose();
        },
        onError: () => {
          alert("Failed to create a tag.");
        },
      },
    );
  }
  return (
    <Overlay click={onClose}>
      <form
        onSubmit={(e) => {
          e.preventDefault();
          createTag();
        }}
        className="m-auto rounded-xl bg-primary-lighter p-4 flex flex-col w-md"
      >
        <input
          type="text"
          placeholder="Tag Name"
          value={tagName}
          onChange={(e) => setTagName(e.target.value)}
          required
          className="bg-primary-lightest p-2 text-xl rounded-md mb-3"
        />
        <div className="flex place-content-between">
          <button
            type="submit"
            className="cursor-pointer rounded-md p-1 border-2 border-primary-light hover:bg-primary-light flex-1 font-bold"
          >
            Add
          </button>
        </div>
      </form>
    </Overlay>
  );
}

function getPageTitle(tagType: TagType) {
  switch (tagType) {
    case "stock":
      return "Preference-StockTags";
    case "transaction":
      return "Preference-TransactionTags";
    case "flow":
      return "Preference-FlowTags";
  }
}
