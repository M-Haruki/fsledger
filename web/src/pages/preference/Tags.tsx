type TagType = "stock" | "transaction" | "flow";

export default function PreferenceTags({ tagType }: { tagType: TagType }) {
  return <>Tags/{tagType}</>;
}
