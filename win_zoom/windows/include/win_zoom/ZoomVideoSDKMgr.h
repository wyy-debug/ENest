#ifndef FLUTTER_PLUGIN_ZOOMVIDEOSDKMGR_H_
#define FLUTTER_PLUGIN_ZOOMVIDEOSDKMGR_H_
#pragma once
#include <windows.h>
#include <vector>
#include <tchar.h>
#include "include/zoom/zoom_video_sdk_delegate_interface.h"
#include "include/zoom/zoom_video_sdk_interface.h"
#include "include/zoom/zoom_video_sdk_session_info_interface.h"
#include "include/helpers/zoom_video_sdk_audio_helper_interface.h"
#include "include/helpers/zoom_video_sdk_video_helper_interface.h"
#include "include/helpers/zoom_video_sdk_share_helper_interface.h"
#include "include/helpers/zoom_video_sdk_chat_helper_interface.h"
#include "include/helpers/zoom_video_sdk_cmd_channel_interface.h"
#include "include/zoom/zoom_video_sdk_api.h"

using namespace ZOOM_VIDEO_SDK_NAMESPACE;
class ZoomVideoSDKMgr
{
public:
    static ZoomVideoSDKMgr& GetInst();
    ~ZoomVideoSDKMgr();

    bool Init(IZoomVideoSDKDelegate* listener, ZoomVideoSDKInitParams init_params);

    void UnInit();
    std::wstring GetErrorStringByErrorCode(ZoomVideoSDKErrors err);

    IZoomVideoSDKSession* JoinSession(ZoomVideoSDKSessionContext& session_context);
    ZoomVideoSDKErrors LeaveSession(bool end);
    ZoomVideoSDKErrors MuteAudio();
    ZoomVideoSDKErrors UnmuteAudio();
    ZoomVideoSDKErrors MuteVideo();
    ZoomVideoSDKErrors UnmuteVideo();
    bool IsMyselfVideoOn();
    bool IsMyselfAudioMuted();
    
    bool SelectCamera(const zchar_t* camera_device_id);	
    ZoomVideoSDKErrors SelectSpeaker(const zchar_t* device_id, const zchar_t* device_name);
    ZoomVideoSDKErrors SelectMic(const zchar_t* device_id, const zchar_t* device_name);
    ZoomVideoSDKErrors StartShareScreen(const zchar_t* monitorID, ZoomVideoSDKShareOption option);
    ZoomVideoSDKErrors StartShareView(HWND hwnd, ZoomVideoSDKShareOption option);
    ZoomVideoSDKErrors StopShare();
    ZoomVideoSDKErrors StartShare2ndCamera();
    std::wstring GetSharingCameraID();
    void SetSharingCameraID(std::wstring cameraID);
    void SetRemoteControlCameraUser(IZoomVideoSDKUser* user);
    IZoomVideoSDKUser* GetRemoteControlCameraUser();
    ZoomVideoSDKErrors RequestControlRemoteCamera();
    ZoomVideoSDKErrors GiveUpControlRemoteCamera(std::wstring& user_name);
    ZoomVideoSDKErrors TurnLeft(unsigned int range);
    ZoomVideoSDKErrors TurnRight(unsigned int range);
    ZoomVideoSDKErrors TurnUp(unsigned int range);
    ZoomVideoSDKErrors TurnDown(unsigned int range);
    ZoomVideoSDKErrors ZoomIn(unsigned int range);
    ZoomVideoSDKErrors ZoomOut(unsigned int range);
    bool RemoveUser(IZoomVideoSDKUser* user);
    void OnUserLeave(IVideoSDKVector<IZoomVideoSDKUser*>* userList);
    
    ZoomVideoSDKErrors SendCommand(IZoomVideoSDKUser* receiver, const zchar_t* strCmd);
    ZoomVideoSDKErrors SendChatToAll(const zchar_t* msgContent);

    const zchar_t* GetSessionName() const;
    int GetUserCountInSession();
    bool IsInSession();
    bool IsHost();
    IZoomVideoSDKUser* GetSharingUser();
    IZoomVideoSDKUser* GetMySelf();
    IZoomVideoSDKUser* GetSessionHost();
    
    IVideoSDKVector<IZoomVideoSDKCameraDevice*>* GetCameraList();
    IVideoSDKVector<IZoomVideoSDKSpeakerDevice*>* GetSpeakerList();
    IVideoSDKVector<IZoomVideoSDKMicDevice*>* GetMicList();
    std::vector<IZoomVideoSDKUser*> GetAllUsers();
    uint32_t GetNumberOfCameras();
    IZoomVideoSDKShareHelper* GetShareHelper();

    ZoomVideoSDKErrors GetSessionAudioStatisticInfo(ZoomVideoSDKSessionAudioStatisticInfo& send_info, ZoomVideoSDKSessionAudioStatisticInfo& recv_info);
    ZoomVideoSDKErrors GetSessionVideoStatisticInfo(ZoomVideoSDKSessionASVStatisticInfo& send_info, ZoomVideoSDKSessionASVStatisticInfo& recv_info);
    ZoomVideoSDKErrors GetSessionShareStatisticInfo(ZoomVideoSDKSessionASVStatisticInfo& send_info, ZoomVideoSDKSessionASVStatisticInfo& recv_info);

private:
    ZoomVideoSDKMgr();

    bool is_inited_;
    
    IZoomVideoSDK* video_sdk_obj_;

    std::wstring sharing_camera_id;
    IZoomVideoSDKUser* remote_camera_control_user;
};
#endif