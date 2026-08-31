export default function Overlay({
  click,
  children,
}: {
  click: () => void;
  children: React.ReactNode;
}) {
  return (
    <div
      onClick={(e) => {
        if (e.target === e.currentTarget) {
          click();
        }
      }}
      className="z-10 w-full h-full bg-gray-500/50 absolute inset-0 flex"
    >
      {children}
    </div>
  );
}
