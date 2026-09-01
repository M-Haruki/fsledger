import type { UUID } from "./uuid";

export interface Flow {
  from: UUID;
  to: UUID;
  fromAmount: bigint;
  toAmount: bigint;
  tags: UUID[];
}
