import { Link } from "react-router-dom";
import { Card } from "@/components/ui/card";

const pages = [
  { to: "/cheers", emoji: "🎉", title: "Daily Cheers", desc: "Send kudos & rotating motivation for your team." },
  { to: "/wins", emoji: "🏆", title: "Team Wins", desc: "Celebrate shipped work and recent victories.", soon: true },
  { to: "/streaks", emoji: "🔥", title: "Streaks", desc: "Track green-CI and clean-test streaks.", soon: true },
  { to: "/thanks", emoji: "💌", title: "Thank-You Notes", desc: "Leave a quick note of appreciation.", soon: true },
  { to: "/breathe", emoji: "🌿", title: "Breathe", desc: "60-second mindful pause.", soon: true },
];

const Index = () => {
  return (
    <main className="min-h-screen bg-background text-foreground">
      <div className="container mx-auto max-w-4xl px-6 py-16">
        <header className="mb-12 text-center">
          <h1 className="text-4xl font-bold tracking-tight">Co-worker Encouragement Hub</h1>
          <p className="mt-3 text-muted-foreground">A small set of pages to lift the team — one daily ritual at a time.</p>
        </header>
        <section className="grid gap-4 md:grid-cols-2">
          {pages.map((p) => (
            <Link key={p.to} to={p.soon ? "#" : p.to} aria-disabled={p.soon}>
              <Card className={`p-6 transition-all hover:shadow-md hover:-translate-y-0.5 ${p.soon ? "opacity-60 cursor-not-allowed" : ""}`}>
                <div className="flex items-start gap-4">
                  <span className="text-3xl" aria-hidden>{p.emoji}</span>
                  <div>
                    <h2 className="text-lg font-semibold">{p.title} {p.soon && <span className="ml-2 text-xs font-normal text-muted-foreground">(coming next)</span>}</h2>
                    <p className="mt-1 text-sm text-muted-foreground">{p.desc}</p>
                  </div>
                </div>
              </Card>
            </Link>
          ))}
        </section>
      </div>
    </main>
  );
};

export default Index;
