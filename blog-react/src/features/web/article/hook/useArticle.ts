import { useCallback, useEffect, useState } from "react";
import { articleInfoByID, type Article } from "@/api/article";

export const useArticle = (id?: string) => {
  const [article, setArticle] = useState<Article | null>(null);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  const load = useCallback(async () => {
    if (!id) return;

    try {
      setLoading(true);
      setError(null);
      const res = await articleInfoByID(id);
      if (res && res?.data && res.code === 0) {
        setArticle(res.data);
      } else {
        setError("load article failed");
      }
    } catch {
      setError("load article failed");
    } finally {
      setLoading(false);
    }
  }, [id]);

  useEffect(() => {
    load();
  }, [load]);

  return {
    article,
    loading,
    error,
    reload: load,
  };
};
