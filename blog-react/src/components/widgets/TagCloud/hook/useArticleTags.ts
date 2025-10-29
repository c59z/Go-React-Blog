import { useEffect, useState } from "react";
import { articleTags } from "@/api/article";
import type { ArticleTag } from "@/api/article";

export const useArticleTags = () => {
  const [tags, setTags] = useState<ArticleTag[]>([]);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<unknown>(null);

  useEffect(() => {
    let alive = true;

    const load = async () => {
      try {
        setLoading(true);
        const res = await articleTags();
        if (alive) {
          setTags(res.data || []);
        }
      } catch (e) {
        if (alive) {
          setError(e);
        }
      } finally {
        if (alive) {
          setLoading(false);
        }
      }
    };

    load();

    return () => {
      alive = false;
    };
  }, []);

  return {
    tags,
    loading,
    error,
  };
};
