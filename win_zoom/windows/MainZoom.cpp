#include "include/win_zoom/MainZoom.h"


MainZoom& MainZoom::GetInstance()
{
	static MainZoom s_MainFramInstance;
	return s_MainFramInstance;
}


void MainZoom::InitVideoSDK()
{
	ZoomVideoSDKInitParams init_params;
	init_params.domain = _T("https://go.zoom.us");	
	init_params.enableLog = true;
	init_params.logFilePrefix = _T("zoom_win_video_demo");
	init_params.videoRawDataMemoryMode = ZoomVideoSDKRawDataMemoryModeHeap;
	init_params.shareRawDataMemoryMode = ZoomVideoSDKRawDataMemoryModeHeap;
	init_params.audioRawDataMemoryMode = ZoomVideoSDKRawDataMemoryModeHeap;
	init_params.enableIndirectRawdata = false;
	MainZoom zoom = GetInstance();
	ZoomVideoSDKMgr::GetInst().Init(&zoom, init_params);
}


void MainZoom::JoinSession()
{

	ZoomVideoSDKSessionContext sessionContext;
	sessionContext.sessionName = L"Session name";
	sessionContext.sessionPassword = L"Session password";
	sessionContext.userName = L"User name";

	// JWT for this session.
	sessionContext.token = L"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhcHBfa2V5IjoiSEppSnRJMjAwWEIxcDRhQ3NKV1h5Y1VzNWJMZDR4Q0NjYW9TIiwicm9sZV90eXBlIjoxLCJ0cGMiOiJ0ZXN0IiwiaWF0IjoxNzQyMjg4NTY3LCJleHAiOjE3NDIyOTIxNjd9.WzuedvmklGMU11Cm48HF3rCF295xFTH9sPpnPv8flrM";

	// Video and audio options.
	sessionContext.videoOption.localVideoOn = true;
	sessionContext.audioOption.connect = true;
	sessionContext.audioOption.mute = false;
	IZoomVideoSDKSession* pSession = ZoomVideoSDKMgr::GetInst().JoinSession(sessionContext);
	std::cout << pSession << std::endl;
}

void MainZoom::onSessionJoin(){}
void MainZoom::onSessionLeave(){}
void MainZoom::onError(ZoomVideoSDKErrors errorCode, int detailErrorCode){}
void MainZoom::onUserJoin(IZoomVideoSDKUserHelper* pUserHelper, IVideoSDKVector<IZoomVideoSDKUser*>* userList){}
void MainZoom::onUserLeave(IZoomVideoSDKUserHelper* pUserHelper, IVideoSDKVector<IZoomVideoSDKUser*>* userList){}


void MainZoom::onUserVideoStatusChanged(IZoomVideoSDKVideoHelper* pVideoHelper, IVideoSDKVector<IZoomVideoSDKUser*>* userList){}
void MainZoom::onUserAudioStatusChanged(IZoomVideoSDKAudioHelper* pAudioHelper, IVideoSDKVector<IZoomVideoSDKUser*>* userList){}
void MainZoom::onUserShareStatusChanged(IZoomVideoSDKShareHelper* pShareHelper, IZoomVideoSDKUser* pUser, IZoomVideoSDKShareAction* pShareAction){}
void MainZoom::onLiveStreamStatusChanged(IZoomVideoSDKLiveStreamHelper* pLiveStreamHelper, ZoomVideoSDKLiveStreamStatus status){}
void MainZoom::onChatNewMessageNotify(IZoomVideoSDKChatHelper* pChatHelper, IZoomVideoSDKChatMessage* messageItem){}
void MainZoom::onUserHostChanged(IZoomVideoSDKUserHelper* pUserHelper, IZoomVideoSDKUser* pUser){}
void MainZoom::onUserActiveAudioChanged(IZoomVideoSDKAudioHelper* pAudioHelper, IVideoSDKVector<IZoomVideoSDKUser*>* list){}

void MainZoom::onSessionNeedPassword(IZoomVideoSDKPasswordHandler* handler){}
void MainZoom::onSessionPasswordWrong(IZoomVideoSDKPasswordHandler* handler){}

void MainZoom::onMixedAudioRawDataReceived(AudioRawData* data_){}
void MainZoom::onOneWayAudioRawDataReceived(AudioRawData* data_, IZoomVideoSDKUser* pUser){}
void MainZoom::onSharedAudioRawDataReceived(AudioRawData* data_){}
void MainZoom::onUserManagerChanged(IZoomVideoSDKUser* pUser){}
void MainZoom::onUserNameChanged(IZoomVideoSDKUser* pUser){}
void MainZoom::onCameraControlRequestResult(IZoomVideoSDKUser* pUser, bool isApproved){}
void MainZoom::onCameraControlRequestReceived(IZoomVideoSDKUser* pUser, ZoomVideoSDKCameraControlRequestType requestType, IZoomVideoSDKCameraControlRequestHandler* pCameraControlRequestHandler){}
void MainZoom::onCommandReceived(IZoomVideoSDKUser* sender, const zchar_t* strCmd){}
void MainZoom::onCommandChannelConnectResult(bool isSuccess){}
void MainZoom::onInviteByPhoneStatus(PhoneStatus status, PhoneFailedReason reason){}
void MainZoom::onCalloutJoinSuccess(IZoomVideoSDKUser* pUser, const zchar_t* phoneNumber){}
void MainZoom::onCloudRecordingStatus(RecordingStatus status, IZoomVideoSDKRecordingConsentHandler* pHandler){}
void MainZoom::onHostAskUnmute(){}
void MainZoom::onMultiCameraStreamStatusChanged(ZoomVideoSDKMultiCameraStreamStatus status, IZoomVideoSDKUser* pUser, IZoomVideoSDKRawDataPipe* pVideoPipe){}



void MainZoom::onMicSpeakerVolumeChanged(unsigned int micVolume, unsigned int speakerVolume)
{
    std::cout << "Mic Volume: " << micVolume << ", Speaker Volume: " << speakerVolume << std::endl;
}
void MainZoom::onAudioDeviceStatusChanged(ZoomVideoSDKAudioDeviceType type, ZoomVideoSDKAudioDeviceStatus status){std::cout<< "123" << std::endl;} 
void MainZoom::onTestMicStatusChanged(ZoomVideoSDK_TESTMIC_STATUS status){std::cout<< "123" << std::endl;} 
void MainZoom::onSelectedAudioDeviceChanged(){std::cout<< "123" << std::endl;} 
 
void MainZoom::onLiveTranscriptionStatus(ZoomVideoSDKLiveTranscriptionStatus status){std::cout<< "123" << std::endl;} 
void MainZoom::onLiveTranscriptionMsgReceived(const zchar_t* ltMsg, IZoomVideoSDKUser* pUser, ZoomVideoSDKLiveTranscriptionOperationType type){std::cout<< "123" << std::endl;} 
void MainZoom::onLiveTranscriptionMsgError(ILiveTranscriptionLanguage* spokenLanguage, ILiveTranscriptionLanguage* transcriptLanguage){std::cout<< "123" << std::endl;} 
