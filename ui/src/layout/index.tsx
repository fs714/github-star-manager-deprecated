import { Layout, Menu, Tooltip } from "antd";
import { ProfileOutlined, BookOutlined, SettingOutlined, TranslationOutlined, ToolTwoTone } from "@ant-design/icons"
import { Link, Outlet, useLocation, useNavigate } from "react-router-dom";
import { useEffect, useState } from "react";
import { ReactImg, GithubImg } from "../assets";
import { useTranslation } from "react-i18next";
import ChangeLanguage from "../components/changeLanguage";


const { Header, Content, Footer } = Layout

export default function AppLayout(): JSX.Element {
    const navigate = useNavigate()
    const location = useLocation()
    const [selectedKey, setSelectedKey] = useState<string[]>([])

    const { t } = useTranslation()

    useEffect(() => {
        if (location.pathname === "/") {
            setSelectedKey(["/repositories"])
        } else {
            setSelectedKey([location.pathname])
        }
    }, [location])

    const onSelect = (e: any) => {
        if (e?.key) {
            navigate(e.key)
        }
    }

    return (
        <Layout className="bg-gray-100">
            <Header className="flex justify-between h-12 leading-[3rem] px-6">
                <div className="flex justify-start">
                    <div className="pr-6 py-2">
                        <Link to="/">
                            <img src={ReactImg} />
                        </Link>
                    </div>
                    <div>
                        <Menu
                            theme="dark"
                            mode="horizontal"
                            selectedKeys={selectedKey}
                            onSelect={onSelect}
                            className="font-bold h-12"
                            items={[
                                {
                                    key: "/repositories",
                                    label: t("layout.REPOSITORY"),
                                    icon: <ProfileOutlined />
                                },
                                {
                                    key: "/tags",
                                    label: t("layout.TAGS"),
                                    icon: <BookOutlined />
                                }
                            ]}
                        />
                    </div>
                </div>
                <div className="flex justify-end">
                    <div>
                        <Menu
                            theme="dark"
                            mode="horizontal"
                            selectedKeys={selectedKey}
                            onSelect={onSelect}
                            className="font-bold h-12"
                            items={[
                                {
                                    key: "/settings",
                                    label: t("layout.SETTINGS"),
                                    icon: <SettingOutlined />
                                }
                            ]}
                        />
                    </div>
                    <div>
                        <ChangeLanguage />
                    </div>
                    <div className="py-2">
                        <Tooltip title="Github">
                            <Link to="https://github.com/" target="_blank">
                                <img src={GithubImg} className="h-8 w-8" />
                            </Link>
                        </Tooltip>
                    </div>
                </div>
            </Header>
            <Content className="bg-white m-3 min-h-[80vh]">
                <Outlet />
            </Content>
            <Footer className="text-center">
                FS714 Design ©2023 Created by FS714
            </Footer>
        </Layout>
    )
}
