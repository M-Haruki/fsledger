import type { DateString } from "./date";
import type { Flow } from "./flow";
import type { UUID } from "./uuid";

export interface Transaction {
  description: string;
  occurredAt: DateString; // example 2026-08-07
  tags: UUID[];
  flows: Flow[];
}
