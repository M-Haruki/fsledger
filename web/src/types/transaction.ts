import type { Flow } from "./flow";
import type { UUID } from "./uuid";

export interface Transaction {
  description: string;
  occurredAt: string; // example 2026-08-07
  tags: UUID[];
  flows: Flow[];
}
