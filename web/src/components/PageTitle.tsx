export default function PageTitle({
  title,
  className,
}: {
  title: string;
  className?: string;
}) {
  return <h1 className={`text-3xl font-bold mb-5 ${className}`}>{title}</h1>;
}
