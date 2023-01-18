import React, { useCallback, useEffect, useMemo, useState } from 'react'
import {
  Container,
  FlexExpander,
  ButtonVariation,
  Layout,
  Text,
  StringSubstitute,
  Button,
  PageError
} from '@harness/uicore'
import * as Diff2Html from 'diff2html'
import cx from 'classnames'
import { useGet } from 'restful-react'
import { useStrings } from 'framework/strings'
import type { GitInfoProps } from 'utils/GitUtils'
import { formatNumber, getErrorMessage, voidFn } from 'utils/Utils'
import { DiffViewer } from 'components/DiffViewer/DiffViewer'
import { useEventListener } from 'hooks/useEventListener'
import { UserPreference, useUserPreference } from 'hooks/useUserPreference'
import { PipeSeparator } from 'components/PipeSeparator/PipeSeparator'
import type { DiffFileEntry } from 'utils/types'
import { DIFF2HTML_CONFIG, ViewStyle } from 'components/DiffViewer/DiffViewerUtils'
import { NoResultCard } from 'components/NoResultCard/NoResultCard'
import { ContainerSpinner } from 'components/ContainerSpinner/ContainerSpinner'
import type { TypesPullReq } from 'services/code'
import { ChangesDropdown } from './ChangesDropdown'
import { DiffViewConfiguration } from './DiffViewConfiguration'
import { ReviewDecisionButton } from './ReviewDecisionButton/ReviewDecisionButton'
import css from './Changes.module.scss'

const STICKY_TOP_POSITION = 64
const STICKY_HEADER_HEIGHT = 150
const diffViewerId = (collection: Unknown[]) => collection.filter(Boolean).join('::::')

interface ChangesProps extends Pick<GitInfoProps, 'repoMetadata'> {
  targetBranch?: string
  sourceBranch?: string
  readOnly?: boolean
  pullRequestMetadata?: TypesPullReq
  emptyTitle: string
  emptyMessage: string
  className?: string
}

export const Changes: React.FC<ChangesProps> = ({
  repoMetadata,
  targetBranch,
  sourceBranch,
  readOnly,
  pullRequestMetadata,
  emptyTitle,
  emptyMessage,
  className
}) => {
  const { getString } = useStrings()
  const [viewStyle, setViewStyle] = useUserPreference(UserPreference.DIFF_VIEW_STYLE, ViewStyle.SIDE_BY_SIDE)
  const [lineBreaks, setLineBreaks] = useUserPreference(UserPreference.DIFF_LINE_BREAKS, false)
  const [diffs, setDiffs] = useState<DiffFileEntry[]>([])
  const [isSticky, setSticky] = useState(false)
  const {
    data: rawDiff,
    error,
    loading,
    refetch
  } = useGet<string>({
    path: `/api/v1/repos/${repoMetadata?.path}/+/compare/${targetBranch}...${sourceBranch}`,
    lazy: !targetBranch || !sourceBranch
  })

  const diffStats = useMemo(
    () =>
      (diffs || []).reduce(
        (obj, diff) => {
          obj.addedLines += diff.addedLines
          obj.deletedLines += diff.deletedLines
          return obj
        },
        { addedLines: 0, deletedLines: 0 }
      ),
    [diffs]
  )

  useEffect(() => {
    const _raw = rawDiff && typeof rawDiff === 'string' ? rawDiff : ''

    if (rawDiff) {
      setDiffs(
        Diff2Html.parse(_raw, DIFF2HTML_CONFIG).map(diff => {
          const viewerId = diffViewerId([diff.oldName, diff.newName])
          const containerId = `container-${viewerId}`
          const contentId = `content-${viewerId}`

          return { ...diff, containerId, contentId }
        })
      )
    }
  }, [rawDiff])

  useEventListener(
    'scroll',
    useCallback(() => setSticky(window.scrollY >= STICKY_HEADER_HEIGHT), [])
  )

  return (
    <Container className={cx(css.container, className)} {...(!!loading || !!error ? { flex: true } : {})}>
      {loading && <ContainerSpinner />}
      {error && <PageError message={getErrorMessage(error)} onClick={voidFn(refetch)} />}
      {!loading &&
        !error &&
        (diffs?.length ? (
          <>
            <Container className={css.header}>
              <Layout.Horizontal>
                <Container flex={{ alignItems: 'center' }}>
                  {/* Files Changed stats */}
                  <Text flex className={css.diffStatsLabel}>
                    <StringSubstitute
                      str={getString('pr.diffStatsLabel')}
                      vars={{
                        changedFilesLink: <ChangesDropdown diffs={diffs} />,
                        addedLines: formatNumber(diffStats.addedLines),
                        deletedLines: formatNumber(diffStats.deletedLines),
                        configuration: (
                          <DiffViewConfiguration
                            viewStyle={viewStyle}
                            setViewStyle={setViewStyle}
                            lineBreaks={lineBreaks}
                            setLineBreaks={setLineBreaks}
                          />
                        )
                      }}
                    />
                  </Text>

                  {/* Show "Scroll to top" button */}
                  {isSticky && (
                    <Layout.Horizontal padding={{ left: 'small' }}>
                      <PipeSeparator height={10} />
                      <Button
                        variation={ButtonVariation.ICON}
                        icon="arrow-up"
                        iconProps={{ size: 14 }}
                        onClick={() => window.scroll({ top: 0 })}
                        tooltip={getString('scrollToTop')}
                        tooltipProps={{ isDark: true }}
                      />
                    </Layout.Horizontal>
                  )}
                </Container>
                <FlexExpander />

                <ReviewDecisionButton
                  repoMetadata={repoMetadata}
                  pullRequestMetadata={pullRequestMetadata}
                  shouldHide={readOnly || pullRequestMetadata?.state === 'merged'}
                />
              </Layout.Horizontal>
            </Container>

            <Layout.Vertical spacing="large" className={cx(css.main, { [css.enableDiffLineBreaks]: lineBreaks })}>
              {diffs?.map((diff, index) => (
                // Note: `key={viewStyle + index + lineBreaks}` resets DiffView when view configuration
                // is changed. Making it easier to control states inside DiffView itself, as it does not
                //  have to deal with any view configuration
                <DiffViewer
                  readOnly={readOnly}
                  index={index}
                  key={viewStyle + index + lineBreaks}
                  diff={diff}
                  viewStyle={viewStyle}
                  stickyTopPosition={STICKY_TOP_POSITION}
                />
              ))}
            </Layout.Vertical>
          </>
        ) : (
          diffs?.length === 0 && (
            <Container padding="xlarge">
              <NoResultCard
                showWhen={() => diffs?.length === 0}
                forSearch={true}
                title={emptyTitle}
                emptySearchMessage={emptyMessage}
              />
            </Container>
          )
        ))}
    </Container>
  )
}
