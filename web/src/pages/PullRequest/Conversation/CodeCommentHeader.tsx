import { ButtonSize, Container, Layout, Text } from '@harnessio/uicore'
import { CodeIcon, type GitInfoProps } from 'utils/GitUtils'
import { CopyButton } from 'components/CopyButton/CopyButton'
interface CodeCommentHeaderProps extends Pick<GitInfoProps, 'repoMetadata' | 'pullReqMetadata'> {
  pullReqMetadata
        `--- a/src/${get(commentItems[0], 'payload.code_comment.path')}`,
        `+++ b/dest/${get(commentItems[0], 'payload.code_comment.path')}`,
          <Layout.Horizontal flex={{ alignItems: 'center', justifyContent: 'flex-start' }}>
            <Text
              inline
              className={css.fname}
              lineClamp={1}
              tooltipProps={{
                portalClassName: css.popover
              }}>
              <Link
                // className={css.fname}
                to={`${routes.toCODEPullRequest({
                  repoPath: repoMetadata?.path as string,
                  pullRequestId: String(pullReqMetadata?.number),
                  pullRequestSection: PullRequestSection.FILES_CHANGED
                })}?path=${commentItems[0].payload?.code_comment?.path}&commentId=${commentItems[0].payload?.id}`}>
                {commentItems[0].payload?.code_comment?.path}
              </Link>
            </Text>
            {commentItems[0].payload?.code_comment?.path && (
              <CopyButton
                content={commentItems[0].payload?.code_comment?.path}
                icon={CodeIcon.Copy}
                size={ButtonSize.MEDIUM}
              />
            )}
          </Layout.Horizontal>