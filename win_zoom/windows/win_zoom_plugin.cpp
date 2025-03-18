#include <windows.h>

#include "win_zoom_plugin.h"
#include "include/zoom_video_sdk_api.h"
#include "include/zoom_video_sdk_interface.h"
#include "include/zoom_video_sdk_def.h"
#include <tchar.h>

// For getPlatformVersion; remove unless needed for your plugin implementation.
#include <VersionHelpers.h>

#include <flutter/method_channel.h>
#include <flutter/plugin_registrar_windows.h>
#include <flutter/standard_method_codec.h>

#include <memory>
#include <sstream>


USING_ZOOM_VIDEO_SDK_NAMESPACE

namespace win_zoom {


    ZoomVideoSDKMgr::ZoomVideoSDKMgr()
    : video_sdk_obj_(nullptr)
    , is_inited_(false)
    , remote_camera_control_user(nullptr)
    {

    }

    ZoomVideoSDKMgr::~ZoomVideoSDKMgr()
    {
      remote_camera_control_user = nullptr;
    }

    ZoomVideoSDKMgr& ZoomVideoSDKMgr::GetInst()
    {
      static ZoomVideoSDKMgr inst;
      return inst;
    }

    bool ZoomVideoSDKMgr::Init(IZoomVideoSDKDelegate* listener, ZoomVideoSDKInitParams init_params)
    {
      if (!is_inited_)
      {
        video_sdk_obj_ = CreateZoomVideoSDKObj();

        if (video_sdk_obj_)
        {
          ZoomVideoSDKErrors err = video_sdk_obj_->initialize(init_params);

          if (ZoomVideoSDKErrors_Success == err)
          {
            is_inited_ = true;
            video_sdk_obj_->addListener(listener);
          }
        }
      }
      return is_inited_;
    }

    void ZoomVideoSDKMgr::UnInit()
    {
      if (is_inited_)
      {
        if (video_sdk_obj_)
        {
          video_sdk_obj_->cleanup();
          DestroyZoomVideoSDKObj();
          video_sdk_obj_ = nullptr;
        }
        is_inited_ = false;
      }
    }

    IZoomVideoSDKSession* ZoomVideoSDKMgr::JoinSession(ZoomVideoSDKSessionContext& session_context)
    {
      IZoomVideoSDKSession* session = nullptr;

      if (video_sdk_obj_)
      {
        session = video_sdk_obj_->joinSession(session_context);
      }

      return session;
    }

    ZoomVideoSDKErrors ZoomVideoSDKMgr::LeaveSession(bool end)
    {
      if (video_sdk_obj_)
      {
        return video_sdk_obj_->leaveSession(end);
      }

      return ZoomVideoSDKErrors_Uninitialize;
    }

    ZoomVideoSDKErrors ZoomVideoSDKMgr::MuteAudio()
    {
      if (video_sdk_obj_)
      {
        IZoomVideoSDKAudioHelper* audio_helper = video_sdk_obj_->getAudioHelper();
        if (audio_helper)
        {
          IZoomVideoSDKUser* my_self = GetMySelf();
          if (my_self)
          {
            return audio_helper->muteAudio(my_self);
          }
        }
      }
      return ZoomVideoSDKErrors_Uninitialize;
    }

    ZoomVideoSDKErrors ZoomVideoSDKMgr::UnmuteAudio()
    {
      if (video_sdk_obj_)
      {
        IZoomVideoSDKAudioHelper* audio_helper = video_sdk_obj_->getAudioHelper();
        if (audio_helper)
        {
          IZoomVideoSDKUser* my_self = GetMySelf();
          if (my_self)
          {
            return audio_helper->unMuteAudio(my_self);
          }
        }
      }
      return ZoomVideoSDKErrors_Uninitialize;
    }

    ZoomVideoSDKErrors ZoomVideoSDKMgr::MuteVideo()
    {
      if (video_sdk_obj_)
      {
        IZoomVideoSDKVideoHelper* video_helper = video_sdk_obj_->getVideoHelper();
        if (video_helper)
        {
          return video_helper->stopVideo();
        }
      }
      return ZoomVideoSDKErrors_Uninitialize;
    }

    ZoomVideoSDKErrors ZoomVideoSDKMgr::UnmuteVideo()
    {
      if (video_sdk_obj_)
      {
        IZoomVideoSDKVideoHelper* video_helper = video_sdk_obj_->getVideoHelper();
        if (video_helper)
        {
          return video_helper->startVideo();
        }
      }
      return ZoomVideoSDKErrors_Uninitialize;
    }

    bool ZoomVideoSDKMgr::IsMyselfVideoOn()
    {
      if (video_sdk_obj_)
      {
        IZoomVideoSDKUser* myself = GetMySelf();
        if (myself)
        {
          if (!myself->GetVideoPipe())
            return false;
          ZoomVideoSDKVideoStatus video_status = myself->GetVideoPipe()->getVideoStatus();
          if (video_status.isOn)
            return true;
          else
            return false;
        }
      }
      return false;
    }

    bool ZoomVideoSDKMgr::IsMyselfAudioMuted()
    {
      if (video_sdk_obj_)
      {
        IZoomVideoSDKUser* myself = GetMySelf();
        if (myself)
        {
          ZoomVideoSDKAudioStatus audio_status = myself->getAudioStatus();
          if (audio_status.isMuted)
            return true;
          else
            return false;
        }
      }
      return false;
    }

    bool ZoomVideoSDKMgr::SelectCamera(const zchar_t* camera_device_id)
    {
      if (video_sdk_obj_)
      {
        IZoomVideoSDKVideoHelper* video_helper = video_sdk_obj_->getVideoHelper();
        if (video_helper)
        {
          return video_helper->selectCamera(camera_device_id);
        }
      }
      return false;
    }

    ZoomVideoSDKErrors ZoomVideoSDKMgr::SelectSpeaker(const zchar_t* device_id, const zchar_t* device_name)
    {
      if (video_sdk_obj_)
      {
        IZoomVideoSDKAudioHelper* audio_helper = video_sdk_obj_->getAudioHelper();
        if (audio_helper)
        {
          return audio_helper->selectSpeaker(device_id, device_name);
        }
      }
      return ZoomVideoSDKErrors_Uninitialize;
    }

    ZoomVideoSDKErrors ZoomVideoSDKMgr::SelectMic(const zchar_t* device_id, const zchar_t* device_name)
    {
      if (video_sdk_obj_)
      {
        IZoomVideoSDKAudioHelper* audio_helper = video_sdk_obj_->getAudioHelper();
        if (audio_helper)
        {
          return audio_helper->selectMic(device_id, device_name);
        }
      }
      return ZoomVideoSDKErrors_Uninitialize;
    }

    ZoomVideoSDKErrors ZoomVideoSDKMgr::StartShareScreen(const zchar_t* monitorID, ZoomVideoSDKShareOption option)
    {
      if (video_sdk_obj_)
      {
        IZoomVideoSDKShareHelper* share_helper = video_sdk_obj_->getShareHelper();
        if (share_helper)
        {
          return share_helper->startShareScreen(monitorID, option);
        }
      }
      return ZoomVideoSDKErrors_Uninitialize;
    }

    ZoomVideoSDKErrors ZoomVideoSDKMgr::StartShareView(HWND hwnd, ZoomVideoSDKShareOption option)
    {
      if (video_sdk_obj_)
      {
        IZoomVideoSDKShareHelper* share_helper = video_sdk_obj_->getShareHelper();
        if (share_helper)
        {
          return share_helper->startShareView(hwnd, option);
        }
      }
      return ZoomVideoSDKErrors_Uninitialize;
    }

    ZoomVideoSDKErrors ZoomVideoSDKMgr::StopShare()
    {
      if (video_sdk_obj_)
      {
        IZoomVideoSDKShareHelper* share_helper = video_sdk_obj_->getShareHelper();
        if (share_helper)
        {
          return share_helper->stopShare();
        }
      }
      return ZoomVideoSDKErrors_Uninitialize;
    }

    ZoomVideoSDKErrors ZoomVideoSDKMgr::SendCommand(IZoomVideoSDKUser* receiver, const zchar_t* strCmd)
    {
      if (video_sdk_obj_)
      {
        IZoomVideoSDKCmdChannel* cmd_channel = video_sdk_obj_->getCmdChannel();
        if (cmd_channel)
        {
          return cmd_channel->sendCommand(receiver, strCmd);
        }
      }
      return ZoomVideoSDKErrors_Uninitialize;
    }

    ZoomVideoSDKErrors ZoomVideoSDKMgr::SendChatToAll(const zchar_t* msgContent)
    {
      if (video_sdk_obj_)
      {
        IZoomVideoSDKChatHelper* chat_helper = video_sdk_obj_->getChatHelper();
        if (chat_helper)
        {
          return chat_helper->sendChatToAll(msgContent);
        }
      }
      return ZoomVideoSDKErrors_Uninitialize;
    }

    const zchar_t* ZoomVideoSDKMgr::GetSessionName() const
    {
      if (video_sdk_obj_)
      {
        IZoomVideoSDKSession* session = video_sdk_obj_->getSessionInfo();
        if (session)
        {
          return session->getSessionName();
        }
      }
      return nullptr;
    }

    int ZoomVideoSDKMgr::GetUserCountInSession()
    {
      if (video_sdk_obj_)
      {
        return GetAllUsers().size();
      }
      return 0;
    }

    bool ZoomVideoSDKMgr::IsInSession()
    {
      if (video_sdk_obj_)
      {
        return video_sdk_obj_->isInSession();
      }
      return false;
    }
    
    bool ZoomVideoSDKMgr::IsHost()
    {
      if (video_sdk_obj_)
      {
        IZoomVideoSDKUser* my_self = GetMySelf();
        if (my_self)
        {
          return my_self->isHost();
        }
      }
      return false;
    }
    
    IZoomVideoSDKUser* ZoomVideoSDKMgr::GetSharingUser()
    {
      if (video_sdk_obj_)
      {
        std::vector<IZoomVideoSDKUser*> user_list = GetAllUsers();
    
        for (size_t i = 0; i < user_list.size(); i++)
        {
          IZoomVideoSDKUser* user = user_list[i];
          if (!user) continue;
    
          IVideoSDKVector<IZoomVideoSDKShareAction*>* pShareActionList = user->getShareActionList();
          if (pShareActionList && pShareActionList->GetCount() > 0)
          {
            IZoomVideoSDKShareAction* pShareAction = pShareActionList->GetItem(0);
            if (pShareAction && pShareAction->getSharePipe())
            {
              ZoomVideoSDKShareStatus share_status = pShareAction->getSharePipe()->getShareStatus();
              if (share_status == ZoomVideoSDKShareStatus_Start || share_status == ZoomVideoSDKShareStatus_Resume)
              {
                return user;
              }
            }
          }
        }
      }
      return nullptr;
    }
    
    IZoomVideoSDKUser* ZoomVideoSDKMgr::GetMySelf()
    {
      IZoomVideoSDKUser* my_self = nullptr;
      if (video_sdk_obj_)
      {
        IZoomVideoSDKSession* session = video_sdk_obj_->getSessionInfo();
        if (session)
        {
          my_self = session->getMyself();
        }
      }
      return my_self;
    }
    
    IZoomVideoSDKUser* ZoomVideoSDKMgr::GetSessionHost()
    {
      if (video_sdk_obj_)
      {
        IZoomVideoSDKSession* session = video_sdk_obj_->getSessionInfo();
        if (session)
        {
          return session->getSessionHost();
        }
      }
      return NULL;
    }
    
    IVideoSDKVector<IZoomVideoSDKCameraDevice*>* ZoomVideoSDKMgr::GetCameraList()
    {
      if (video_sdk_obj_)
      {
        IZoomVideoSDKVideoHelper* video_helper = video_sdk_obj_->getVideoHelper();
        if (video_helper)
        {
          return video_helper->getCameraList();
        }
      }
      return nullptr;
    }
    
    uint32_t ZoomVideoSDKMgr::GetNumberOfCameras()
    {
      if (video_sdk_obj_)
      {
        IZoomVideoSDKVideoHelper* video_helper = video_sdk_obj_->getVideoHelper();
        if (video_helper)
        {
          return video_helper->getNumberOfCameras();
        }
      }
      return 0;
    }

    ZoomVideoSDKErrors ZoomVideoSDKMgr::StartShare2ndCamera()
    {
      ZoomVideoSDKErrors err = ZoomVideoSDKErrors_Uninitialize;
      if (video_sdk_obj_)
      {
        IZoomVideoSDKCameraDevice* second_camera = NULL;
        IZoomVideoSDKShareHelper* share_helper = video_sdk_obj_->getShareHelper();
        IVideoSDKVector<IZoomVideoSDKCameraDevice*>* camera_list = GetCameraList();
        if (share_helper &&
          camera_list &&
          camera_list->GetCount() > 1
          )
        {
          for (size_t i = 0; i < camera_list->GetCount(); i++)
          {
            IZoomVideoSDKCameraDevice* camera_temp = camera_list->GetItem(i);
            if (camera_temp && !camera_temp->isSelectedDevice())
            {
              err = share_helper->startShare2ndCamera(camera_temp->getDeviceId());
              if (err == ZoomVideoSDKErrors_Success)
              {
                SetSharingCameraID(camera_temp->getDeviceId());
                break;
              }
            }
          }
        }
      }
    
      return err;
    }

    IZoomVideoSDKShareHelper* ZoomVideoSDKMgr::GetShareHelper()
    {
      if (video_sdk_obj_)
      {
        IZoomVideoSDKShareHelper* share_helper = video_sdk_obj_->getShareHelper();
        if (share_helper)
        {
          return share_helper;
        }
      }
      return NULL;
    }

    std::wstring ZoomVideoSDKMgr::GetSharingCameraID()
    {
      return sharing_camera_id;
    }

    void ZoomVideoSDKMgr::SetSharingCameraID(std::wstring cameraID)
    {
      sharing_camera_id = cameraID;
    }

    void ZoomVideoSDKMgr::SetRemoteControlCameraUser(IZoomVideoSDKUser* user)
    {
      remote_camera_control_user = user;
    }

    IZoomVideoSDKUser* ZoomVideoSDKMgr::GetRemoteControlCameraUser()
    {
      return remote_camera_control_user;
    }

    ZoomVideoSDKErrors ZoomVideoSDKMgr::RequestControlRemoteCamera()
    {
      ZoomVideoSDKErrors err = ZoomVideoSDKErrors_Internal_Error;
      if (remote_camera_control_user && remote_camera_control_user->getRemoteCameraControlHelper())
      {
        err = remote_camera_control_user->getRemoteCameraControlHelper()->requestControlRemoteCamera();
      }
      return err;
    }

    ZoomVideoSDKErrors ZoomVideoSDKMgr::GiveUpControlRemoteCamera(std::wstring& user_name)
    {
      ZoomVideoSDKErrors err = ZoomVideoSDKErrors_Internal_Error;
      if (!remote_camera_control_user)
      {
        return err;
      }
      if (remote_camera_control_user && remote_camera_control_user->getRemoteCameraControlHelper())
      {
        err = remote_camera_control_user->getRemoteCameraControlHelper()->giveUpControlRemoteCamera();
      }
      if (err == ZoomVideoSDKErrors_Success &&
        remote_camera_control_user->getUserName())
      {
        user_name = remote_camera_control_user->getUserName();
      }
      return err;
    }

    ZoomVideoSDKErrors ZoomVideoSDKMgr::ZoomIn(unsigned int range)
    {
      ZoomVideoSDKErrors err = ZoomVideoSDKErrors_Internal_Error;
      if (remote_camera_control_user && remote_camera_control_user->getRemoteCameraControlHelper())
      {
        err = remote_camera_control_user->getRemoteCameraControlHelper()->zoomIn(range);
      }
      return err;
    }
    
    ZoomVideoSDKErrors ZoomVideoSDKMgr::ZoomOut(unsigned int range)
    {
      ZoomVideoSDKErrors err = ZoomVideoSDKErrors_Internal_Error;
      if (remote_camera_control_user && remote_camera_control_user->getRemoteCameraControlHelper())
      {
        err = remote_camera_control_user->getRemoteCameraControlHelper()->zoomOut(range);
      }
      return err;
    }
    
    ZoomVideoSDKErrors ZoomVideoSDKMgr::TurnDown(unsigned int range)
    {
      ZoomVideoSDKErrors err = ZoomVideoSDKErrors_Internal_Error;
      if (remote_camera_control_user && remote_camera_control_user->getRemoteCameraControlHelper())
      {
        err = remote_camera_control_user->getRemoteCameraControlHelper()->turnDown(range);
      }
      return err;
    }

    ZoomVideoSDKErrors ZoomVideoSDKMgr::TurnUp(unsigned int range)
    {
      ZoomVideoSDKErrors err = ZoomVideoSDKErrors_Internal_Error;
      if (remote_camera_control_user && remote_camera_control_user->getRemoteCameraControlHelper())
      {
        err = remote_camera_control_user->getRemoteCameraControlHelper()->turnUp(range);
      }
      return err;
    }

    ZoomVideoSDKErrors ZoomVideoSDKMgr::TurnLeft(unsigned int range)
    {
      ZoomVideoSDKErrors err = ZoomVideoSDKErrors_Internal_Error;
      if (remote_camera_control_user && remote_camera_control_user->getRemoteCameraControlHelper())
      {
        err = remote_camera_control_user->getRemoteCameraControlHelper()->turnLeft(range);
      }
      return err;
    }

    ZoomVideoSDKErrors ZoomVideoSDKMgr::TurnRight(unsigned int range)
    {
      ZoomVideoSDKErrors err = ZoomVideoSDKErrors_Internal_Error;
      if (remote_camera_control_user && remote_camera_control_user->getRemoteCameraControlHelper())
      {
        err = remote_camera_control_user->getRemoteCameraControlHelper()->turnRight(range);
      }
      return err;
    }

    void ZoomVideoSDKMgr::OnUserLeave(IVideoSDKVector<IZoomVideoSDKUser*>* userList)
    {
      if (!userList || !remote_camera_control_user) return;
    
      std::wstring curUserId = _T("");
      if (remote_camera_control_user && remote_camera_control_user->getUserID())
      {
        curUserId = remote_camera_control_user->getUserID();
      }
    
      for (int i = 0; i < userList->GetCount(); i++)
      {
        IZoomVideoSDKUser* leave_user = userList->GetItem(i);
        if (!leave_user) continue;
    
        std::wstring userId = leave_user->getUserID();
        if (curUserId == userId)
        {
          SetRemoteControlCameraUser(NULL);
        }
      }
    }

    bool ZoomVideoSDKMgr::RemoveUser(IZoomVideoSDKUser* user)
    {
      bool bret = false;
      if (video_sdk_obj_)
      {
        IZoomVideoSDKUserHelper* pUserHelper = video_sdk_obj_->getUserHelper();
        if (pUserHelper)
        {
          bret = pUserHelper->removeUser(user);
        }
      }
      return bret;
    }
    
    IVideoSDKVector<IZoomVideoSDKSpeakerDevice*>* ZoomVideoSDKMgr::GetSpeakerList()
    {
      if (video_sdk_obj_)
      {
        IZoomVideoSDKAudioHelper* audio_helper = video_sdk_obj_->getAudioHelper();
        if (audio_helper)
        {
          return audio_helper->getSpeakerList();
        }
      }
      return nullptr;
    }
    
    IVideoSDKVector<IZoomVideoSDKMicDevice*>* ZoomVideoSDKMgr::GetMicList()
    {
      if (video_sdk_obj_)
      {
        IZoomVideoSDKAudioHelper* audio_helper = video_sdk_obj_->getAudioHelper();
        if (audio_helper)
        {
          return audio_helper->getMicList();
        }
      }
      return nullptr;
    }
    
    std::vector<IZoomVideoSDKUser*> ZoomVideoSDKMgr::GetAllUsers()
    {
      std::vector<IZoomVideoSDKUser*> vect_all_user;
    
      if (video_sdk_obj_)
      {
        IZoomVideoSDKSession* session = video_sdk_obj_->getSessionInfo();
        if (session)
        {
          IZoomVideoSDKUser* my_self = session->getMyself();
          if (my_self)
          {
            vect_all_user.push_back(my_self);
          }
    
          IVideoSDKVector<IZoomVideoSDKUser*>* remote_users = session->getRemoteUsers();
    
          if (remote_users)
          {
            for (int i = 0; i < remote_users->GetCount(); i++)
            {
              IZoomVideoSDKUser* user = remote_users->GetItem(i);
              if (!user) continue;
              vect_all_user.push_back(user);
            }
          }
        }
      }
      return vect_all_user;
    }
    
    ZoomVideoSDKErrors ZoomVideoSDKMgr::GetSessionAudioStatisticInfo(ZoomVideoSDKSessionAudioStatisticInfo& send_info, ZoomVideoSDKSessionAudioStatisticInfo& recv_info)
    {
      if (video_sdk_obj_)
      {
        IZoomVideoSDKSession* session = video_sdk_obj_->getSessionInfo();
        if (session)
        {
          session->getSessionAudioStatisticInfo(send_info, recv_info);
        }
      }
    
      return ZoomVideoSDKErrors_Uninitialize;
    }
    
    ZoomVideoSDKErrors ZoomVideoSDKMgr::GetSessionVideoStatisticInfo(ZoomVideoSDKSessionASVStatisticInfo& send_info, ZoomVideoSDKSessionASVStatisticInfo& recv_info)
    {
      if (video_sdk_obj_)
      {
        IZoomVideoSDKSession* session = video_sdk_obj_->getSessionInfo();
        if (session)
        {
          session->getSessionVideoStatisticInfo(send_info, recv_info);
        }
      }
    
      return ZoomVideoSDKErrors_Uninitialize;
    }
    
    ZoomVideoSDKErrors ZoomVideoSDKMgr::GetSessionShareStatisticInfo(ZoomVideoSDKSessionASVStatisticInfo& send_info, ZoomVideoSDKSessionASVStatisticInfo& recv_info)
    {
      if (video_sdk_obj_)
      {
        IZoomVideoSDKSession* session = video_sdk_obj_->getSessionInfo();
        if (session)
        {
          session->getSessionShareStatisticInfo(send_info, recv_info);
        }
      }
    
      return ZoomVideoSDKErrors_Uninitialize;
    }
    
    std::wstring ZoomVideoSDKMgr::GetErrorStringByErrorCode(ZoomVideoSDKErrors err)
    {
      std::wstring sValue;
    
      switch (err)
      {
      case ZoomVideoSDKErrors_Success:
      {
        sValue = _T("ZoomVideoSDKErrors_Success");
        break;
      }
      case ZoomVideoSDKErrors_Wrong_Usage:
      {
        sValue = _T("ZoomVideoSDKErrors_Wrong_Usage");
        break;
      }
      case ZoomVideoSDKErrors_Internal_Error:
      {
        sValue = _T("ZoomVideoSDKErrors_Internal_Error");
        break;
      }
      case ZoomVideoSDKErrors_Uninitialize:
      {
        sValue = _T("ZoomVideoSDKErrors_Uninitialize");
        break;
      }
      case ZoomVideoSDKErrors_Memory_Error:
      {
        sValue = _T("ZoomVideoSDKErrors_Memory_Error");
        break;
      }
      case ZoomVideoSDKErrors_Load_Module_Error:
      {
        sValue = _T("ZoomVideoSDKErrors_Load_Module_Error");
        break;
      }
      case ZoomVideoSDKErrors_UnLoad_Module_Error:
      {
        sValue = _T("ZoomVideoSDKErrors_UnLoad_Module_Error");
        break;
      }
      case ZoomVideoSDKErrors_Invalid_Parameter:
      {
        sValue = _T("ZoomVideoSDKErrors_Invalid_Parameter");
        break;
      }
      case ZoomVideoSDKErrors_Call_Too_Frequently:
      {
        sValue = _T("ZoomVideoSDKErrors_Call_Too_Frequently");
        break;
      }
      case ZoomVideoSDKErrors_Unknown:
      {
        sValue = _T("ZoomVideoSDKErrors_Unknown");
        break;
      }
      case ZoomVideoSDKErrors_Auth_Error:
      {
        sValue = _T("ZoomVideoSDKErrors_Auth_Error");
        break;
      }
      case ZoomVideoSDKErrors_Auth_Empty_Key_or_Secret:
      {
        sValue = _T("ZoomVideoSDKErrors_Auth_Empty_Key_or_Secret");
        break;
      }
      case ZoomVideoSDKErrors_Auth_Wrong_Key_or_Secret:
      {
        sValue = _T("ZoomVideoSDKErrors_Auth_Wrong_Key_or_Secret");
        break;
      }
      case ZoomVideoSDKErrors_Auth_DoesNot_Support_SDK:
      {
        sValue = _T("ZoomVideoSDKErrors_Auth_DoesNot_Support_SDK");
        break;
      }
      case ZoomVideoSDKErrors_Auth_Disable_SDK:
      {
        sValue = _T("ZoomVideoSDKErrors_Auth_Disable_SDK");
        break;
      }
      case ZoomVideoSDKErrors_SessionModule_Not_Found:
      {
        sValue = _T("ZoomVideoSDKErrors_SessionModule_Not_Found");
        break;
      }
      case ZoomVideoSDKErrors_SessionService_Invalid:
      {
        sValue = _T("ZoomVideoSDKErrors_SessionService_Invalid");
        break;
      }
      case ZoomVideoSDKErrors_Session_Join_Failed:
      {
        sValue = _T("ZoomVideoSDKErrors_Session_Join_Failed");
        break;
      }
      case ZoomVideoSDKErrors_Session_No_Rights:
      {
        sValue = _T("ZoomVideoSDKErrors_Session_No_Rights");
        break;
      }
      case ZoomVideoSDKErrors_Session_Already_In_Progress:
      {
        sValue = _T("ZoomVideoSDKErrors_Session_Already_In_Progress");
        break;
      }
      case ZoomVideoSDKErrors_Session_Dont_Support_SessionType:
      {
        sValue = _T("ZoomVideoSDKErrors_Session_Dont_Support_SessionType");
        break;
      }
      case ZoomVideoSDKErrors_Session_Reconnecting:
      {
        sValue = _T("ZoomVideoSDKErrors_Session_Reconnecting");
        break;
      }
      case ZoomVideoSDKErrors_Session_Disconnecting:
      {
        sValue = _T("ZoomVideoSDKErrors_Session_Disconnecting");
        break;
      }
      case ZoomVideoSDKErrors_Session_Not_Started:
      {
        sValue = _T("ZoomVideoSDKErrors_Session_Not_Started");
        break;
      }
      case ZoomVideoSDKErrors_Session_Need_Password:
      {
        sValue = _T("ZoomVideoSDKErrors_Session_Need_Password");
        break;
      }
      case ZoomVideoSDKErrors_Session_Password_Wrong:
      {
        sValue = _T("ZoomVideoSDKErrors_Session_Password_Wrong");
        break;
      }
      case ZoomVideoSDKErrors_Session_Remote_DB_Error:
      {
        sValue = _T("ZoomVideoSDKErrors_Session_Remote_DB_Error");
        break;
      }
      case ZoomVideoSDKErrors_Session_Invalid_Param:
      {
        sValue = _T("ZoomVideoSDKErrors_Session_Invalid_Param");
        break;
      }
      case ZoomVideoSDKErrors_Session_Audio_Error:
      {
        sValue = _T("ZoomVideoSDKErrors_Session_Audio_Error");
        break;
      }
      case ZoomVideoSDKErrors_Session_Audio_No_Microphone:
      {
        sValue = _T("ZoomVideoSDKErrors_Session_Audio_No_Microphone");
        break;
      }
      case ZoomVideoSDKErrors_Session_Video_Error:
      {
        sValue = _T("ZoomVideoSDKErrors_Session_Video_Error");
        break;
      }
      case ZoomVideoSDKErrors_Session_Video_Device_Error:
      {
        sValue = _T("ZoomVideoSDKErrors_Session_Video_Device_Error");
        break;
      }
      case ZoomVideoSDKErrors_Session_Live_Stream_Error:
      {
        sValue = _T("ZoomVideoSDKErrors_Session_Live_Stream_Error");
        break;
      }
      case ZoomVideoSDKErrors_Session_Phone_Error:
      {
        sValue = _T("ZoomVideoSDKErrors_Session_Phone_Error");
        break;
      }
      case ZoomVideoSDKErrors_RAWDATA_MALLOC_FAILED:
      {
        sValue = _T("ZoomVideoSDKErrors_RAWDATA_MALLOC_FAILED");
        break;
      }
      case ZoomVideoSDKErrors_RAWDATA_NOT_IN_Session:
      {
        sValue = _T("ZoomVideoSDKErrors_RAWDATA_NOT_IN_Session");
        break;
      }
      case ZoomVideoSDKErrors_RAWDATA_NO_LICENSE:
      {
        sValue = _T("ZoomVideoSDKErrors_RAWDATA_NO_LICENSE");
        break;
      }
      case ZoomVideoSDKErrors_RAWDATA_VIDEO_MODULE_NOT_READY:
      {
        sValue = _T("ZoomVideoSDKErrors_RAWDATA_VIDEO_MODULE_NOT_READY");
        break;
      }
      case ZoomVideoSDKErrors_RAWDATA_VIDEO_MODULE_ERROR:
      {
        sValue = _T("ZoomVideoSDKErrors_RAWDATA_VIDEO_MODULE_ERROR");
        break;
      }
      case ZoomVideoSDKErrors_RAWDATA_VIDEO_DEVICE_ERROR:
      {
        sValue = _T("ZoomVideoSDKErrors_RAWDATA_VIDEO_DEVICE_ERROR");
        break;
      }
      case ZoomVideoSDKErrors_RAWDATA_NO_VIDEO_DATA:
      {
        sValue = _T("ZoomVideoSDKErrors_RAWDATA_NO_VIDEO_DATA");
        break;
      }
      case ZoomVideoSDKErrors_RAWDATA_SHARE_MODULE_NOT_READY:
      {
        sValue = _T("ZoomVideoSDKErrors_RAWDATA_SHARE_MODULE_NOT_READY");
        break;
      }
      case ZoomVideoSDKErrors_RAWDATA_SHARE_MODULE_ERROR:
      {
        sValue = _T("ZoomVideoSDKErrors_RAWDATA_SHARE_MODULE_ERROR");
        break;
      }
      case ZoomVideoSDKErrors_RAWDATA_AUDIO_MODULE_NOT_READY:
      {
        sValue = _T("ZoomVideoSDKErrors_RAWDATA_AUDIO_MODULE_NOT_READY");
        break;
      }
      case ZoomVideoSDKErrors_RAWDATA_AUDIO_MODULE_ERROR:
      {
        sValue = _T("ZoomVideoSDKErrors_RAWDATA_AUDIO_MODULE_ERROR");
        break;
      }
      case ZoomVideoSDKErrors_RAWDATA_NO_AUDIO_DATA:
      {
        sValue = _T("ZoomVideoSDKErrors_RAWDATA_NO_AUDIO_DATA");
        break;
      }
      case ZoomVideoSDKErrors_RAWDATA_PREPROCESS_RAWDATA_ERROR:
      {
        sValue = _T("ZoomVideoSDKErrors_RAWDATA_PREPROCESS_RAWDATA_ERROR");
        break;
      }
      case ZoomVideoSDKErrors_RAWDATA_NO_DEVICE_RUNNING:
      {
        sValue = _T("ZoomVideoSDKErrors_RAWDATA_NO_DEVICE_RUNNING");
        break;
      }
      case ZoomVideoSDKErrors_RAWDATA_INIT_DEVICE:
      {
        sValue = _T("ZoomVideoSDKErrors_RAWDATA_INIT_DEVICE");
        break;
      }
      case ZoomVideoSDKErrors_RAWDATA_VIRTUAL_DEVICE:
      {
        sValue = _T("ZoomVideoSDKErrors_RAWDATA_VIRTUAL_DEVICE");
        break;
      }
      case ZoomVideoSDKErrors_RAWDATA_CANNOT_CHANGE_VIRTUAL_DEVICE_IN_PREVIEW:
      {
        sValue = _T("ZoomVideoSDKErrors_RAWDATA_CANNOT_CHANGE_VIRTUAL_DEVICE_IN_PREVIEW");
        break;
      }
      case ZoomVideoSDKErrors_RAWDATA_INTERNAL_ERROR:
      {
        sValue = _T("ZoomVideoSDKErrors_RAWDATA_INTERNAL_ERROR");
        break;
      }
      case ZoomVideoSDKErrors_RAWDATA_SEND_TOO_MUCH_DATA_IN_SINGLE_TIME:
      {
        sValue = _T("ZoomVideoSDKErrors_RAWDATA_SEND_TOO_MUCH_DATA_IN_SINGLE_TIME");
        break;
      }
      case ZoomVideoSDKErrors_RAWDATA_SEND_TOO_FREQUENTLY:
      {
        sValue = _T("ZoomVideoSDKErrors_RAWDATA_SEND_TOO_FREQUENTLY");
        break;
      }
      case ZoomVideoSDKErrors_RAWDATA_VIRTUAL_MIC_IS_TERMINATE:
      {
        sValue = _T("ZoomVideoSDKErrors_RAWDATA_VIRTUAL_MIC_IS_TERMINATE");
        break;
      }
      case ZoomVideoSDKErrors_Session_Share_Error:
      {
        sValue = _T("ZoomVideoSDKErrors_Session_Share_Error");
        break;
      }
      case ZoomVideoSDKErrors_Session_Share_Module_Not_Ready:
      {
        sValue = _T("ZoomVideoSDKErrors_Session_Share_Module_Not_Ready");
        break;
      }
      case ZoomVideoSDKErrors_Session_Share_You_Are_Not_Sharing:
      {
        sValue = _T("ZoomVideoSDKErrors_Session_Share_You_Are_Not_Sharing");
        break;
      }
      case ZoomVideoSDKErrors_Session_Share_Type_Is_Not_Support:
      {
        sValue = _T("ZoomVideoSDKErrors_Session_Share_Type_Is_Not_Support");
        break;
      }
      case ZoomVideoSDKErrors_Session_Share_Internal_Error:
      {
        sValue = _T("ZoomVideoSDKErrors_Session_Share_Internal_Error");
        break;
      }
      default:
      {
        sValue = _T("ZoomVideoSDKErrors_UNKNOWN");
        break;
      }
      }
      return sValue;
    }



// API

void WinZoomPlugin::RegisterWithRegistrar(flutter::PluginRegistrarWindows *registrar) {
  auto channel = std::make_unique<flutter::MethodChannel<flutter::EncodableValue>>(
      registrar->messenger(), "win_zoom",
      &flutter::StandardMethodCodec::GetInstance());

  auto plugin = std::make_unique<WinZoomPlugin>();
  channel->SetMethodCallHandler(
      [plugin_pointer = plugin.get()](const auto &call, auto result) {
        plugin_pointer->HandleMethodCall(call, std::move(result));
      });
  registrar->AddPlugin(std::move(plugin));
}

WinZoomPlugin::WinZoomPlugin() {}

WinZoomPlugin::~WinZoomPlugin() 
{
    UninitVideoSDK();
}

void WinZoomPlugin::InitVideoSDK()
{
  ZoomVideoSDKInitParams init_params;
  init_params.domain = _T("https://go.zoom.us");
  init_params.enableLog = true;
  init_params.logFilePrefix = _T("zoom_win_video_demo");
  init_params.videoRawDataMemoryMode = ZoomVideoSDKRawDataMemoryModeHeap;
  init_params.shareRawDataMemoryMode = ZoomVideoSDKRawDataMemoryModeHeap;
  init_params.audioRawDataMemoryMode = ZoomVideoSDKRawDataMemoryModeHeap;
  init_params.enableIndirectRawdata = false;

  ZoomVideoSDKMgr::GetInst().Init(this, init_params);
}

void WinZoomPlugin::UninitVideoSDK()
{
  ZoomVideoSDKMgr::GetInst().UnInit();
}

void WinZoomPlugin::initSDK(std::unique_ptr<flutter::MethodResult<flutter::EncodableValue>> result)
{
  if (video_sdk_ == nullptr) {
    video_sdk_ = CreateZoomVideoSDKObj();
    if (video_sdk_) {
      ZOOM_VIDEO_SDK_NAMESPACE::ZoomVideoSDKInitParams init_params;
      init_params.domain = L"https://zoom.us";
      init_params.enableIndirectRawdata = false;
      init_params.enableLog = true;
      init_params.logFilePrefix = L"prefix";
      init_params.audioRawDataMemoryMode = ZoomVideoSDKRawDataMemoryModeHeap;
      init_params.videoRawDataMemoryMode = ZoomVideoSDKRawDataMemoryModeHeap;
      init_params.shareRawDataMemoryMode = ZoomVideoSDKRawDataMemoryModeHeap;
      
      ZoomVideoSDKErrors err = video_sdk_->initialize(init_params);
      if (err == ZoomVideoSDKErrors_Success) {
        result->Success("init success");
        return;
      }
      else {
        std::string error_message = "初始化失败，错误代码: " + std::to_string(static_cast<int>(err));
        OutputDebugStringA(error_message.c_str());
        
        switch(err) {
          case ZoomVideoSDKErrors_Auth_Error:
            result->Error("AUTH_FAILED", "认证失败");
            break;
          case ZoomVideoSDKErrors_Uninitialize:
            result->Error("CONFIG_ERROR", "SDK配置错误");
            break;
          default:
            result->Error("UNKNOWN_ERROR", error_message);
        }
      }
    }
  }
}

void WinZoomPlugin::joinSession(const flutter::MethodCall<flutter::EncodableValue> &method_call, std::unique_ptr<flutter::MethodResult<flutter::EncodableValue>> result) {
    const auto* arguments = std::get_if<flutter::EncodableMap>(method_call.arguments());
    if (arguments && video_sdk_) {
      auto session_name_it = arguments->find(flutter::EncodableValue("sessionName"));
      auto session_password_it = arguments->find(flutter::EncodableValue("sessionPassword"));
      auto user_name_it = arguments->find(flutter::EncodableValue("userName"));
      
      if (session_name_it != arguments->end() && 
          session_password_it != arguments->end() && 
          user_name_it != arguments->end()) {
        const std::string& session_name = std::get<std::string>(session_name_it->second);
        const std::string& session_password = std::get<std::string>(session_password_it->second);
        const std::string& user_name = std::get<std::string>(user_name_it->second);
        
        ZOOM_VIDEO_SDK_NAMESPACE::ZoomVideoSDKSessionContext session_context;
        std::wstring wsession_name(session_name.begin(), session_name.end());
        std::wstring wsession_password(session_password.begin(), session_password.end());
        std::wstring wuser_name(user_name.begin(), user_name.end());
        session_context.sessionName = wsession_name.c_str();
        session_context.sessionPassword = wsession_password.c_str();
        session_context.userName = wuser_name.c_str();
        session_context.audioOption.connect = true;
        session_context.videoOption.localVideoOn = true;
        
        IZoomVideoSDKSession* session = video_sdk_->joinSession(session_context);
        if (session != nullptr) {
          result->Success(flutter::EncodableValue(true));
          return;
        }
        result->Error("JOIN_FAILED", "Failed to join session");
      } else {
        result->Error("INVALID_ARGUMENT", "Session name, password and user name are required");
      }
    } else {
      result->Error("SDK_NOT_INITIALIZED", "Zoom Video SDK is not initialized");
    }
}

void WinZoomPlugin::leaveSession(std::unique_ptr<flutter::MethodResult<flutter::EncodableValue>> result) {
  if (video_sdk_) {
    ZoomVideoSDKErrors err = video_sdk_->leaveSession(false);
    if (err == ZoomVideoSDKErrors_Success) {
      result->Success(flutter::EncodableValue(true));
      return;
    }
    result->Error("LEAVE_FAILED", "Failed to leave session");
  } else {
    result->Error("SDK_NOT_INITIALIZED", "Zoom Video SDK is not initialized");
  }
}


void WinZoomPlugin::HandleMethodCall(
    const flutter::MethodCall<flutter::EncodableValue> &method_call,
    std::unique_ptr<flutter::MethodResult<flutter::EncodableValue>> result) {
  const std::string& method_name = method_call.method_name();
  
  switch(method_name[0]) {
    case 'i':
      if (method_name == "initSDK") {
        initSDK(std::move(result));
        return;
      }
      break;
    case 'j':
      if (method_name == "joinSession") {
        joinSession(method_call, std::move(result));
        return;
      }
      break;
    case 'l':
      if (method_name == "leaveSession") {
        leaveSession(std::move(result));
        return;
      }
      break;
  }
  
  result->NotImplemented();
}

}  // namespace win_zoom
