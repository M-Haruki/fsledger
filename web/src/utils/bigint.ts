export function bigint2string(value: bigint, exponent: number): string {
  const sign = value < 0n ? "-" : "";
  const str = (value < 0n ? -value : value).toString();
  if (exponent === 0) {
    return sign + str;
  }
  if (exponent > 0) {
    const padded = str.padStart(exponent + 1, "0");
    const pos = padded.length - exponent;
    return `${sign}${padded.slice(0, pos)}.${padded.slice(pos)}`;
  }
  return sign + str + "0".repeat(-exponent);
}

export function string2bigint(value: string, exponent: number): bigint {
  if (!/^[+-]?(?:\d+\.?\d*|\.\d+)$/.test(value)) {
    return 0n;
  }
  const sign = value.startsWith("-") ? -1n : 1n;
  const str = value.replace(/^[+-]/, "");
  const [integer = "", decimal = ""] = str.split(".");
  if (exponent === 0) {
    return sign * BigInt(integer || "0");
  }
  if (exponent > 0) {
    const decimalPart = decimal.padEnd(exponent, "0").slice(0, exponent);
    return (
      sign *
      (BigInt(integer || "0") * 10n ** BigInt(exponent) +
        BigInt(decimalPart || "0"))
    );
  }
  return sign * (BigInt(integer || "0") / 10n ** BigInt(-exponent));
}
