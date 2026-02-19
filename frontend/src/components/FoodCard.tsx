import { useCart } from "@/context/CartContext";
import { Card, CardHeader, CardTitle, CardDescription, CardContent, CardFooter } from "@/components/ui/card";
import { getComboRecommendations } from "@/api/recommendations";
import { useState } from "react";
import { Button } from "@/components/ui/button";
import { Plus } from "lucide-react";
import RecommendationBar from "./RecommendationBar";
import { FoodItem } from "@/types/food";

interface FoodCardProps {
  item : FoodItem
}

const FoodCard = ({ item }: FoodCardProps) => {
  const { addToCart } = useCart();
  const [recommendations, setRecommendations] = useState([]);

  const handleAddToCart = async () => {
    addToCart(item);

    const recos = await getComboRecommendations(item.id);
    setRecommendations(recos);
  };

  return (
    <>
      <Card className="glass-card overflow-hidden group hover-lift border-white/5">
        <div className="h-52 overflow-hidden relative">
          <img
            src={"/assets/" + item.image}
            alt={item.name}
            className="h-full w-full object-cover transition-transform duration-700 group-hover:scale-110"
          />
          <div className="absolute inset-0 bg-gradient-to-t from-background/90 to-transparent opacity-60" />
        </div>

        <CardHeader className="p-4 pb-2">
          <CardTitle className="text-xl font-bold tracking-tight">{item.name}</CardTitle>
          <CardDescription className="text-sm text-foreground/70 line-clamp-2">{item.description}</CardDescription>
        </CardHeader>

        <CardContent className="p-4 pt-1">
          <p className="text-2xl font-black text-accent drop-shadow-[0_0_8px_hsl(var(--accent)/0.5)]">${item.price}</p>
        </CardContent>

        <CardFooter className="p-4 pt-1">
          <Button
            className="w-full gap-2 bg-primary hover:bg-primary/80 h-11 text-sm font-bold neon-glow transition-all active:scale-95"
            onClick={handleAddToCart}
          >
            <Plus className="h-4 w-4" /> Add to Experience
          </Button>
        </CardFooter>
      </Card>

      {recommendations.length > 0 && (
        <RecommendationBar title="Perfect Pairings" items={recommendations} />
      )}
    </>
  );
};

export default FoodCard;
