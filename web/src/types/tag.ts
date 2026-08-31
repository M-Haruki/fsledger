export type TagType = "stock" | "transaction" | "flow";
export interface Tag {
  id: string;
  name: string;
}
export const NullTag: Tag = {
  id: "",
  name: "",
};
