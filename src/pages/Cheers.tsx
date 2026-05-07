import { useEffect, useState } from "react";
import { Link } from "react-router-dom";
import { Card } from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Textarea } from "@/components/ui/textarea";
import { toast } from "sonner";

const QUOTES = [
  "Small commits, big momentum. You're moving the needle. 🚀",
  "Green tests today are the gift to your future self. ✅",
  "Refactor with kindness — your teammates read this code too. 💛",
  "Done > perfect. Ship it, learn, iterate. 🔁",
  "The best PR is the one that's been reviewed. Tag a friend. 👀",
  "You debugged that? Legend. 🕵️",
  "Coverage isn't just a number — it's a hug for tomorrow's you. 🤗",
  "One brave question saves an hour of guessing. 🙋",
];

type Cheer = { id: number; to: string; from: string; msg: string; ts: number };
const STORAGE = "cheers.v1";

const Cheers = () => {
  const [quote, setQuote] = useState(QUOTES[0]);
  const [cheers, setCheers] = useState<Cheer[]>([]);
  const [to, setTo] = useState("");
  const [from, setFrom] = useState("");
  const [msg, setMsg] = useState("");

  useEffect(() => {
    setQuote(QUOTES[Math.floor(Math.random() * QUOTES.length)]);
    try {
      const raw = localStorage.getItem(STORAGE);
      if (raw) setCheers(JSON.parse(raw));
    } catch { /* ignore */ }
  }, []);

  const saveCheer = (e: React.FormEvent) => {
    e.preventDefault();
    if (!to.trim() || !msg.trim()) {
      toast.error("Please add a teammate name and a short message.");
      return;
    }
    const next: Cheer = { id: Date.now(), to: to.trim(), from: from.trim() || "Anonymous", msg: msg.trim(), ts: Date.now() };
    const updated = [next, ...cheers].slice(0, 50);
    setCheers(updated);
    localStorage.setItem(STORAGE, JSON.stringify(updated));
    setTo(""); setMsg("");
    toast.success("Cheer sent! 🎉");
  };

  const newQuote = () => {
    let q = quote;
    while (q === quote && QUOTES.length > 1) q = QUOTES[Math.floor(Math.random() * QUOTES.length)];
    setQuote(q);
  };

  return (
    <main className="min-h-screen bg-background text-foreground">
      <div className="container mx-auto max-w-3xl px-6 py-12">
        <Link to="/" className="text-sm text-muted-foreground hover:text-foreground">← Back to hub</Link>
        <header className="mt-4 mb-10">
          <h1 className="text-3xl font-bold tracking-tight">🎉 Daily Cheers</h1>
          <p className="mt-2 text-muted-foreground">A rotating boost and a wall of kudos for your teammates.</p>
        </header>

        <Card className="p-8 mb-8 bg-secondary/40">
          <p className="text-xl font-medium leading-relaxed">"{quote}"</p>
          <Button variant="outline" size="sm" className="mt-4" onClick={newQuote}>New quote</Button>
        </Card>

        <Card className="p-6 mb-8">
          <h2 className="text-lg font-semibold mb-4">Send a cheer</h2>
          <form onSubmit={saveCheer} className="grid gap-3">
            <Input placeholder="To (teammate name)" value={to} onChange={(e) => setTo(e.target.value)} />
            <Input placeholder="From (your name, optional)" value={from} onChange={(e) => setFrom(e.target.value)} />
            <Textarea placeholder="Your message of appreciation…" value={msg} onChange={(e) => setMsg(e.target.value)} rows={3} />
            <Button type="submit" className="justify-self-start">Send cheer 🎉</Button>
          </form>
        </Card>

        <section>
          <h2 className="text-lg font-semibold mb-4">Recent cheers</h2>
          {cheers.length === 0 ? (
            <p className="text-sm text-muted-foreground">No cheers yet — be the first to brighten someone's day.</p>
          ) : (
            <ul className="space-y-3">
              {cheers.map((c) => (
                <li key={c.id}>
                  <Card className="p-4">
                    <div className="flex items-baseline justify-between gap-3">
                      <p className="text-sm font-medium">To <span className="text-primary">{c.to}</span> · from {c.from}</p>
                      <span className="text-xs text-muted-foreground">{new Date(c.ts).toLocaleString()}</span>
                    </div>
                    <p className="mt-2 text-sm">{c.msg}</p>
                  </Card>
                </li>
              ))}
            </ul>
          )}
        </section>
      </div>
    </main>
  );
};

export default Cheers;
