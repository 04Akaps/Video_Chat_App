import NiceModal, {useModal} from "@ebay/nice-modal-react";
import { Tooltip} from "antd";
import React, {useEffect} from "react";

export default NiceModal.create((props) => {
    const modal = useModal();

    useEffect(() => {
        if (modal.visible) {
            const timer = setTimeout(() => {
                modal.hide();
            }, 1000);

            return () => {
                clearTimeout(timer);
            };
        }
    }, [modal.visible]);

    console.log(props)

    return (
        <Tooltip
            title={props.text}
            visible={modal.visible}
            afterVisibleChange={(visible) => {
                if (!visible) {
                    modal.remove();
                }
            }}
            footer={null}
        >
          ------------------------------
        </Tooltip>
    );
});