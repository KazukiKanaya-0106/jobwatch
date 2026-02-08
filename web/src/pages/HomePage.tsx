import { Button } from "@/components/ui/button";

export default function HomePage() {
  return (
    <>
      <h1 className="font-bold text-3xl underline">Home Page</h1>
      <Button onClick={() => console.log("clicked")}>Click me</Button>
    </>
  );
}
