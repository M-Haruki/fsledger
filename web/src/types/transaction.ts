import type { Flow } from "./flow";

export interface Transaction {
  description: string;
  occurredAt: string; // example 2026-08-07
  tags: string[];
  flows: Flow[];
}
