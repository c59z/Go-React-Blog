import type { CommentItemProps } from "./types";
import "./index.scss";

const CommentItem = ({
  comment,
  level,
  onReply,
  onLike,
  onDelete,
  canDelete,
}: CommentItemProps) => {
  const { id, user, content, created_at } = comment;

  return (
    <div className={`article-comment-item level-${level}`}>
      <img
        className="comment-avatar"
        src={`${import.meta.env.VITE_SERVER_URL}${user.avatar}`}
        alt={user.username}
      />

      <div className="comment-body">
        <div className="comment-header">
          <span className="comment-username">{user.username}</span>
          <span className="comment-time">
            {new Date(created_at).toLocaleString()}
          </span>

          {canDelete && (
            <button
              className="comment-delete"
              onClick={() => onDelete(id.toString())}
            >
              删除
            </button>
          )}
        </div>

        <div className="comment-content">{content}</div>

        <div className="comment-actions">
          <button onClick={() => onLike(id.toString())}>👍 点赞</button>
          <button onClick={() => onReply(id.toString())}>回复</button>
        </div>
      </div>
    </div>
  );
};

export default CommentItem;
