import { useEffect, useState } from "react";
import {
  useListTags,
  useUpdateTag,
  getListTagsQueryKey,
  useDeleteTag,
  useCreateTag,
} from "@/utils/tagWrapper.ts";
import PageTitle from "@/components/PageTitle";
import { type TagType, type Tag, NullTag } from "@/types/tag";
import Overlay from "@/components/Overlay";
import { useQueryClient } from "@tanstack/react-query";

export default function PreferenceTags({ tagType }: { tagType: TagType }) {
  const { data, isPending, isError } = useListTags(tagType);
  const [editTag, setEditTag] = useState<Tag>(NullTag);
  const [isShowEditer, setIsShowEditer] = useState(false);
  const [isShowAdder, setIsShowAdder] = useState(false);
  return (
    <>
      <TagEditer
        tag={editTag}
        tagType={tagType}
        onClose={() => setIsShowEditer(false)}
        enable={isShowEditer}
      />
      <TagAdder
        tagType={tagType}
        onClose={() => setIsShowAdder(false)}
        enable={isShowAdder}
      />
      <PageTitle className="flex place-content-between">
        {getPageTitle(tagType)}
        <AddBtn onAdd={() => setIsShowAdder(true)} />
      </PageTitle>
      {isPending ? (
        <p>Loading</p>
      ) : isError || data?.status != 200 ? (
        <p>Error</p>
      ) : (
        <div className="flex gap-2 flex-wrap">
          {data?.data.map((tag) => (
            <ATag
              tag={{ id: tag.id, name: tag.name }}
              key={tag.id}
              onTagSelect={(tag) => {
                setEditTag(tag);
                setIsShowEditer(true);
              }}
            />
          ))}
        </div>
      )}
    </>
  );
}

function ATag({
  tag,
  onTagSelect,
}: {
  tag: Tag;
  onTagSelect: (tag: Tag) => void;
}) {
  return (
    <div
      onClick={() => onTagSelect(tag)}
      className="border w-xs border-primary-light rounded-xl p-2 flex gap-x-2 text-xl cursor-pointer hover:bg-primary-lighter"
    >
      {tag.name}
    </div>
  );
}

function TagEditer({
  tag,
  tagType,
  onClose,
  enable = true,
}: {
  tag: Tag;
  tagType: TagType;
  onClose: () => void;
  enable: boolean;
}) {
  const [tagName, setTagName] = useState("");
  const changeTagMutation = useUpdateTag(tagType);
  const deleteTagMutation = useDeleteTag(tagType);
  const queryClient = useQueryClient();
  useEffect(() => {
    setTagName(tag.name);
  }, [tag]);
  if (!enable) {
    return;
  }
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
          alert("Failed to change the tag name");
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
          alert("Failed to delete the tag name");
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
        className="m-auto rounded-xl bg-primary-lighter p-4 flex flex-col"
      >
        <input
          type="text"
          value={tagName}
          onChange={(e) => setTagName(e.target.value)}
          required
          className="bg-primary-lightest p-2 text-xl rounded-md mb-3"
        />
        <div className="flex place-content-between">
          <button
            type="submit"
            className="cursor-pointer rounded-md p-1 border-2 border-primary-light hover:bg-primary-light"
          >
            Change
          </button>
          <button
            type="button"
            onClick={deleteTag}
            className="cursor-pointer rounded-md p-1 border-2 border-primary-light hover:bg-primary-light"
          >
            🗑️
          </button>
        </div>
      </form>
    </Overlay>
  );
}

function TagAdder({
  tagType,
  onClose,
  enable = true,
}: {
  tagType: TagType;
  onClose: () => void;
  enable: boolean;
}) {
  const [tagName, setTagName] = useState("");
  const createTagMutation = useCreateTag(tagType);
  const queryClient = useQueryClient();

  if (!enable) {
    return;
  }
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
          alert("Failed to create a tag");
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
        className="m-auto rounded-xl bg-primary-lighter p-4 flex flex-col"
      >
        <input
          type="text"
          value={tagName}
          onChange={(e) => setTagName(e.target.value)}
          required
          className="bg-primary-lightest p-2 text-xl rounded-md mb-3"
        />
        <div className="flex place-content-between">
          <button
            type="submit"
            className="cursor-pointer rounded-md p-1 border-2 border-primary-light hover:bg-primary-light"
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

function AddBtn({ onAdd }: { onAdd: () => void }) {
  return (
    <button
      onClick={onAdd}
      className="text-xl font-medium rounded-xl p-1 boder bg-primary-lighter hover:bg-primary-light"
    >
      +
    </button>
  );
}
