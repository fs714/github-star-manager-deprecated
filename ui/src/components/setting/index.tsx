import { Card, Drawer, type DrawerProps, Radio } from 'antd';
import React from 'react';
import { useTranslation } from 'react-i18next';

import { useThemeContext } from '../../context/theme';

export default function Settings({
  setShowDrawer,
  ...props
}: { setShowDrawer: React.Dispatch<React.SetStateAction<boolean>> } & DrawerProps): JSX.Element {
  const { t } = useTranslation();
  const { theme, setTheme } = useThemeContext();
  const { menuStyle } = theme;

  return (
    <div>
      <Drawer
        title={<div className="text-transparent">Setting</div>}
        closable={false}
        width="400px"
        placement="right"
        {...props}
        headerStyle={{ background: menuStyle === 'light' ? 'white' : 'black', height: '3em' }}
        bodyStyle={{ padding: 0 }}
      >
        <div>
          <Card title={t('layout.THEME_SETTING')}>
            <div className="flex justify-between items-center">
              <div className="font-bold text-base text-gray-600">{t('layout.MENUSTYLE')}</div>
              <Radio.Group
                defaultValue="light"
                buttonStyle="solid"
                onChange={(e) => setTheme({ ...theme, menuStyle: e.target.value })}
              >
                <Radio.Button value="light">{t('layout.LIGHT')}</Radio.Button>
                <Radio.Button value="dark">{t('layout.DARK')}</Radio.Button>
              </Radio.Group>
            </div>
          </Card>
        </div>
      </Drawer>
    </div>
  );
}
