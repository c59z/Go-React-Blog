import { commentDelete, type Comment } from "@/api/comment";
import { useArticleComments } from "./hooks/useArticleComments";
import type { RenderComment } from "./types";
import CommentItem from "./CommentItem";
import "./index.scss";
import CommentEditor from "./CommentEditor";
import { useCreateComment } from "./hooks/useCreateComment";
import { useToast } from "../widgets/Toast";
import { useRequireLogin } from "@/hook/useRequireLogin";
import { canDeleteComment } from "@/utils/role";
import { useUserStore } from "@/stores/user";

interface Props {
  articleId: string;
}

const CommentSection = ({ articleId }: Props) => {
  const { comments, loading, error, reload, deleteComment } =
    useArticleComments(articleId);
  const { create, loading: loadingAddComment } = useCreateComment();
  const userInfo = useUserStore((e) => e.user);
  const requireLogin = useRequireLogin();

  const { toast, Toast } = useToast();

  if (loading) return <div>评论加载中...</div>;
  if (error) return <div>评论加载失败</div>;
  if (!comments.length) return <div>暂无评论</div>;

  const flattenComments = (comments: Comment[]): RenderComment[] => {
    const result: RenderComment[] = [];

    const walk = (list: Comment[], level: 1 | 2 | 3) => {
      for (const item of list) {
        result.push({ comment: item, level });

        if (item.children?.length) {
          const nextLevel = level === 3 ? 3 : ((level + 1) as 2 | 3);
          walk(item.children, nextLevel);
        }
      }
    };

    walk(comments, 1);

    return result;
  };

  const renderList = flattenComments(comments);

  const handleSubmit = async (content: string) => {
    requireLogin(async () => {
      const ok = await create({
        article_id: articleId,
        content: content,
        p_id: null,
      });
      if (ok) {
        toast.success("评论发送成果");
        reload();
      } else {
        toast.error("评论发送失败");
      }
    });
  };

  const handleDelete = async (commentId: number) => {
    const ok = await commentDelete({
      ids: [commentId],
    });
    if (ok) {
      toast.success("评论删除成果");
      deleteComment(commentId);
    } else {
      toast.error("评论删除失败");
    }
  };

  return (
    <>
      <CommentEditor
        loading={loadingAddComment}
        onSubmit={handleSubmit}
      ></CommentEditor>
      <div className="article-comment-list">
        {renderList.map(({ comment, level }) => (
          <CommentItem
            key={comment.id}
            comment={comment}
            level={level}
            canDelete={canDeleteComment(comment, userInfo)}
            onDelete={() => handleDelete(comment.id)}
            onReplySuccess={() => {
              reload();
            }}
          />
        ))}
      </div>
      {Toast}
    </>
  );
};

export default CommentSection;
