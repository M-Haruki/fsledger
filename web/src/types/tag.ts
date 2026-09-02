import { type UUID } from "./uuid";

export type TagType = "stock" | "transaction" | "flow";
export interface Tag {
  id: UUID;
  name: string;
}
export const NullTag: Tag = {
  id: "",
  name: "",
};
