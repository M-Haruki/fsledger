export function Switch({
  value,
  onChange,
  className = "",
}: {
  value: boolean;
  onChange: () => void;
  className?: string;
}) {
  return (
    <div
      onClick={onChange}
      className={`w-18 h-10 p-1 rounded-3xl bg-primary-lightest flex items-center ${className}`}
    >
      <div
        className={`size-8 rounded-3xl ${value ? "bg-primary-light ml-auto" : "bg-primary-lighter mr-auto"}`}
      />
    </div>
  );
}
