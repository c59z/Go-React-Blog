import { useEffect, useState, useCallback } from "react";
import { commentInfoByArticleID, type Comment } from "@/api/comment";

export const useArticleComments = (articleId?: string) => {
  const [comments, setComments] = useState<Comment[]>([]);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  // --- 加载评论 ---
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

  // --- 方法 1：递归删除评论 ---
  const removeCommentById = useCallback(
    (id: number, list: Comment[]): Comment[] => {
      return list
        .filter((c) => c.id !== id)
        .map((c) => ({
          ...c,
          children: c.children ? removeCommentById(id, c.children) : [],
        }));
    },
    []
  );

  const deleteComment = useCallback(
    (id: number) => {
      setComments((prev) => removeCommentById(id, prev));
    },
    [removeCommentById]
  );

  const updateComments = useCallback(
    (updater: (prev: Comment[]) => Comment[]) => {
      setComments((prev) => updater(prev));
    },
    []
  );

  return {
    comments,
    loading,
    error,
    reload: load,
    deleteComment,
    updateComments,
  };
};
