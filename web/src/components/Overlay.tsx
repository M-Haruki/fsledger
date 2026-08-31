export function Overlay({
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

export function OverlayLoading() {
  return (
    <div className="z-10 w-full h-full bg-gray-500/50 absolute inset-0 flex flex-col justify-center items-center">
      <div className="h-15 w-15 animate-spin rounded-full border-4 border-primary-main border-t-transparent mb-5" />
      <p className="text-4xl">Loading...</p>
    </div>
  );
}
