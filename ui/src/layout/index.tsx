import { Layout, Menu, Tooltip } from "antd";
import { ProfileOutlined, BookOutlined, SettingOutlined } from "@ant-design/icons"
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
        <Layout className="relative flex h-screen w-screen flex-col bg-gray-100">
            <Header className="flex h-12 w-screen items-center justify-between px-6 leading-[3rem]">
                <div className="flex justify-start">
                    <div className="flex items-center bg-inherit">
                        <Link to="/">
                            <img src={ReactImg} className="flex" />
                        </Link>
                    </div>
                    <div>
                        <Menu
                            theme="dark"
                            mode="horizontal"
                            selectedKeys={selectedKey}
                            onSelect={onSelect}
                            className="h-12 bg-inherit px-6 font-bold"
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
                <div className="flex justify-start">
                    <div>
                        <Menu
                            theme="dark"
                            mode="horizontal"
                            selectedKeys={selectedKey}
                            onSelect={onSelect}
                            className="h-12 bg-inherit pr-3 font-bold"
                            items={[
                                {
                                    key: "/settings",
                                    label: t("layout.SETTINGS"),
                                    icon: <SettingOutlined />
                                }
                            ]}
                        />
                    </div>
                    <div className="bg-inherit pr-6">
                        <ChangeLanguage />
                    </div>
                    <div className="flex items-center bg-inherit">
                        <Tooltip title="Github">
                            <Link to="https://github.com/" target="_blank">
                                <img src={GithubImg} className="flex h-8 w-8" />
                            </Link>
                        </Tooltip>
                    </div>
                </div>
            </Header>
            <Content className="grow bg-white m-3">
                <Outlet />
            </Content>
            <Footer className="text-center mb-2 p-1 font-bold text-black/[.45]">
                Copyright © 2023 Fs714 Labs
            </Footer>
        </Layout>
    )
}
