export interface Stock {
  name: string;
  hasAmount: boolean;
  currency: string;
  currencyExponent: number;
  description: string;
  tags: string[];
}
