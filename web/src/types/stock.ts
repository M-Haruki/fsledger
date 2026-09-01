export interface Stock {
  name: string;
  hasAmount: boolean;
  currency: string;
  currencyExponent: number;
  description: string;
  tags: string[];
}

export interface StockAbstract {
  id: string;
  name: string;
  hasAmount: boolean;
  currency: string;
  currencyExponent: number;
}
