import { useState, useCallback } from "react";
import { commentCreate } from "@/api/comment";
import type { CommentCreateRequest } from "@/api/comment";

export const useCreateComment = () => {
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  const create = useCallback(async (data: CommentCreateRequest) => {
    try {
      setLoading(true);
      setError(null);

      const res = await commentCreate(data);

      if (res.code !== 0) {
        throw new Error("create comment failed");
      }

      return true;
    } catch (a) {
      setError("create comment failed");
      console.log(a);
      return false;
    } finally {
      setLoading(false);
    }
  }, []);

  return {
    create,
    loading,
    error,
  };
};
