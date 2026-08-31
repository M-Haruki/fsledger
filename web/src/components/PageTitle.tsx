export default function PageTitle({
  className,
  children,
}: {
  className?: string;
  children: React.ReactNode;
}) {
  return <h1 className={`text-3xl font-bold mb-5 ${className}`}>{children}</h1>;
}
