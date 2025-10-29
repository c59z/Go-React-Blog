import { commentNew, type CommentVM } from "@/api/comment";
import dayjs from "dayjs";
import { useEffect, useState } from "react";

export const useCommentNew = () => {
  const [comments, setComments] = useState<CommentVM[]>([]);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<unknown>(null);

  useEffect(() => {
    let alive = true;

    const load = async () => {
      try {
        setLoading(true);
        const res = await commentNew();
        if (!alive) return;

        const list: CommentVM[] = (res.data || []).map((item) => ({
          ...item,
          created_at_text: dayjs(item.created_at).fromNow(),
        }));

        setComments(list);
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
    comments,
    loading,
    error,
  };
};
