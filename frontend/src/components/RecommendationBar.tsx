import { Card, CardContent } from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { useCart } from "@/context/CartContext";

const RecommendationBar = ({ title, items }) => {
  const { addToCart } = useCart();

  if (!items || items.length === 0) return null;

  return (
    <div className="mt-6 p-6 glass-card rounded-2xl animate-float">
      <h3 className="text-xl font-black mb-4 text-neon">Glow-Up Your Meal</h3>

      <div className="flex gap-6 overflow-x-auto no-scrollbar pb-2">
        {items.map((item) => (
          <Card
            key={item.id}
            className="w-56 min-w-[14rem] glass-card border-white/5 hover-lift"
          >
            <div className="relative">
              <img
                src={"/assets/" + item.image}
                alt={item.name}
                className="h-32 w-full object-cover rounded-t-xl"
              />
              <div className="absolute inset-0 bg-gradient-to-t from-background/80 to-transparent" />
            </div>

            <CardContent className="p-4 space-y-2">
              <p className="font-bold text-base line-clamp-1">{item.name}</p>
              <p className="text-xs text-muted-foreground line-clamp-2">{item.description}</p>
              <p className="text-accent font-black text-lg mt-1">${item.price}</p>

              <Button
                className="w-full h-10 mt-3 text-xs font-bold bg-primary/20 hover:bg-primary text-foreground border border-primary/50 transition-all"
                onClick={() => addToCart(item)}
              >
                Upgrade +
              </Button>
            </CardContent>
          </Card>
        ))}
      </div>
    </div>
  );
};

export default RecommendationBar;
