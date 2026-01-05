import { useEffect, useState, useCallback } from "react";
import { commentInfoByArticleID, type Comment } from "@/api/comment";

export const useArticleComments = (articleId?: string) => {
  const [comments, setComments] = useState<Comment[]>([]);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  const load = useCallback(async () => {
    if (!articleId) return;

    try {
      setLoading(true);
      setError(null);

      const res = await commentInfoByArticleID(articleId);

      if (res.code === 0 && Array.isArray(res.data)) {
        setComments(res.data);
      } else {
        setError("load comments failed");
      }
    } catch {
      setError("load comments failed");
    } finally {
      setLoading(false);
    }
  }, [articleId]);

  useEffect(() => {
    load();
  }, [load]);

  return {
    comments,
    loading,
    error,
    reload: load,
  };
};
