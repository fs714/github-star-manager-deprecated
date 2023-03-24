import React from 'react';
import { Card, Drawer, type DrawerProps } from 'antd';
import { useTranslation } from 'react-i18next';

export default function Settings({ setShowDrawer, ...props }:
  { setShowDrawer: React.Dispatch<React.SetStateAction<boolean>> } & DrawerProps): JSX.Element {
  const { t } = useTranslation()

  return (
    <div>
      <Drawer
        title="Setting"
        closable={false}
        width="400px"
        placement="right" {...props}
        headerStyle={{ background: 'black', height: '3em' }}
        bodyStyle={{ padding: 0 }}
      >
        <div>
          <Card title={t("layout.THEME_SETTING")}>
            <p>Some contents...</p>
            <p>Some contents...</p>
            <p>Some contents...</p>
          </Card>
        </div>
      </Drawer>
    </div>
  )
}
