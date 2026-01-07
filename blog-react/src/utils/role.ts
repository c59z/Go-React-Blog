import type { Comment } from "@/api/comment";
import type { User } from "@/api/user";

export const canDeleteComment = (
  comment: Comment,
  currentUser: User | null
): boolean => {
  if (!currentUser) return false;

  if (currentUser.role_id === 2) return true;

  if (comment.user_uuid === currentUser.uuid) return true;

  return false;
};
