import { uuid4 } from '@sentry/utils';
import { useDispatch } from 'react-redux';
import { useObservable } from 'react-use';

import { Correlation, DataSourceRef } from '@grafana/data';
import { getDataSourceSrv } from '@grafana/runtime';

import { updateDataSource } from '../datasources/state/actions';

export const useCorrelations = () => {
  const dispatch = useDispatch();

  const datasources = useObservable(getDataSourceSrv().$datasources) || [];

  const correlations = datasources.filter((ds) => ds.correlations.length >= 1) || [];

  const add = (sourceUid: string, correlation: Omit<Correlation, 'uid'>) => {
    const sourceDs = datasources.find((ds) => ds.uid === sourceUid)!;
    dispatch(
      updateDataSource({
        ...sourceDs,
        correlations: [...sourceDs.correlations, { ...correlation, uid: uuid4() }],
      })
    );
  };

  const remove = (source: DataSourceRef, correlationUid: Correlation['uid']) => {
    dispatch(
      updateDataSource({
        ...source,
        correlations: source.correlations.filter((c) => c.uid !== correlationUid),
      })
    );
  };

  return { correlations, add, remove };
};
