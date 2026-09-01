export interface Flow {
  from: string;
  to: string;
  fromAmount: bigint;
  toAmount: bigint;
  tags: string[];
}
