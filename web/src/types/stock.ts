import type { UUID } from "./uuid";

export interface Stock {
  name: string;
  hasAmount: boolean;
  currency: string;
  currencyExponent: number;
  description: string;
  tags: UUID[];
}

export interface StockAbstract {
  id: UUID;
  name: string;
  hasAmount: boolean;
  currency: string;
  currencyExponent: number;
}
