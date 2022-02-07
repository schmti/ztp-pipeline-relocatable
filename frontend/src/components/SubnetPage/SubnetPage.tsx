import React from 'react';

import { Page } from '../Page';
import { ContentThreeRows } from '../ContentThreeRows';
import { WizardProgress } from '../WizardProgress';
import { useWizardProgressContext } from '../WizardProgress/WizardProgressContext';
import { WizardFooter } from '../WizardFooter';

export const SubnetPage: React.FC = () => {
  const { setActiveStep } = useWizardProgressContext();
  React.useEffect(() => setActiveStep('subnet'), [setActiveStep]);

  return (
    <Page>
      <ContentThreeRows
        top={<WizardProgress />}
        middle={<div>TODO: https://marvelapp.com/prototype/hfd719b/screen/84707949/handoff</div>}
        bottom={<WizardFooter back={undefined} next="virtualip" />}
      />
    </Page>
  );
};
