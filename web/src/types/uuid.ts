export type UUID = string & { readonly __brand: "uuid" };

export const NullUUID: UUID = parseUUID("00000000-0000-0000-0000-000000000000");

export function parseUUID(id: string): UUID {
  return id as UUID;
}
