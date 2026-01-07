import type { Comment } from "@/api/comment";

export interface CommentItemProps {
  comment: Comment;
  level: 1 | 2 | 3;
  canDelete: boolean;

  /** 事件回调 */
  onDelete: (commentId: string) => void;
  onReplySuccess: () => void;
}
