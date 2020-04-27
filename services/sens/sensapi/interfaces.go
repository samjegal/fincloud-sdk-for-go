package sensapi

// FINCLOUD_APACHE_NO_VERSION

import (
    "context"
    "github.com/samjegal/fincloud-sdk-for-go/services/sens"
    "github.com/Azure/go-autorest/autorest"
)

        // SMSClientAPI contains the set of methods on the SMSClient type.
        type SMSClientAPI interface {
            DeleteReservation(ctx context.Context, serviceID string, reserveID string) (result autorest.Response, err error)
            DeleteSchedule(ctx context.Context, serviceID string, scheduleCode string, messageID string) (result autorest.Response, err error)
            GetRequest(ctx context.Context, serviceID string, requestID string) (result sens.SMSMessageParameter, err error)
            GetResult(ctx context.Context, serviceID string, messageID string) (result sens.SMSMessageParameter, err error)
            Request(ctx context.Context, serviceID string, parameter sens.SMSMessageRequestParameter) (result sens.SMSMessageParameter, err error)
        }

        var _ SMSClientAPI = (*sens.SMSClient)(nil)
        // ChannelClientAPI contains the set of methods on the ChannelClient type.
        type ChannelClientAPI interface {
            AddUser(ctx context.Context, serviceID string, channelName string, userID string) (result autorest.Response, err error)
            Create(ctx context.Context, serviceID string, channel sens.PushChannelRequestParameter) (result autorest.Response, err error)
            DeleteUser(ctx context.Context, serviceID string, channelName string, userID string) (result autorest.Response, err error)
        }

        var _ ChannelClientAPI = (*sens.ChannelClient)(nil)
        // MessageClientAPI contains the set of methods on the MessageClient type.
        type MessageClientAPI interface {
            Delete(ctx context.Context, serviceID string, reserveID string) (result autorest.Response, err error)
            GetResult(ctx context.Context, serviceID string, requestID string) (result sens.PushMessageResultResponseParameter, err error)
            Send(ctx context.Context, serviceID string, parameter sens.PushMessageRequestParameter) (result sens.PushMessageResponseParameter, err error)
        }

        var _ MessageClientAPI = (*sens.MessageClient)(nil)
        // ScheduleClientAPI contains the set of methods on the ScheduleClient type.
        type ScheduleClientAPI interface {
            Create(ctx context.Context, serviceID string, parameter sens.PushScheduleRegisterParameter) (result autorest.Response, err error)
            Delete(ctx context.Context, serviceID string, scheduleCode string) (result autorest.Response, err error)
            Get(ctx context.Context, serviceID string, scheduleCode string) (result sens.PushScheduleFetchParameter, err error)
            List(ctx context.Context, serviceID string, page *int32, scheduleCode string, size *int32, sortParameter string) (result sens.PushScheduleFetchAllParameter, err error)
            Put(ctx context.Context, serviceID string, scheduleCode string, parameter sens.PushScheduleRegisterParameter) (result sens.PushScheduleFetchParameter, err error)
        }

        var _ ScheduleClientAPI = (*sens.ScheduleClient)(nil)
        // ScheduleMessageClientAPI contains the set of methods on the ScheduleMessageClient type.
        type ScheduleMessageClientAPI interface {
            Delete(ctx context.Context, serviceID string, scheduleCode string, messageID string) (result autorest.Response, err error)
            Get(ctx context.Context, serviceID string, scheduleCode string, messageID string) (result sens.PushScheduleFetchParameter, err error)
            List(ctx context.Context, serviceID string, scheduleCode string, page *int32, size *int32) (result sens.PushScheduleFetchAllParameter, err error)
        }

        var _ ScheduleMessageClientAPI = (*sens.ScheduleMessageClient)(nil)
        // DeviceClientAPI contains the set of methods on the DeviceClient type.
        type DeviceClientAPI interface {
            Create(ctx context.Context, serviceID string, parameter sens.PushUserRequestParameter) (result autorest.Response, err error)
            Delete(ctx context.Context, serviceID string, userID string) (result autorest.Response, err error)
            Get(ctx context.Context, serviceID string, userID string) (result sens.PushUserResponseParameter, err error)
        }

        var _ DeviceClientAPI = (*sens.DeviceClient)(nil)
        // AlimtalkClientAPI contains the set of methods on the AlimtalkClient type.
        type AlimtalkClientAPI interface {
            Create(ctx context.Context, serviceID string, request sens.AlimTalkMessageRequestParameter) (result sens.AlimTalkMessageResponseParameter, err error)
            DeleteReservation(ctx context.Context, serviceID string, reserveID string) (result autorest.Response, err error)
            DeleteSchedule(ctx context.Context, serviceID string, scheduleCode string, messageID string) (result autorest.Response, err error)
            GetRequest(ctx context.Context, serviceID string, requestID string) (result sens.AlimTalkRequestResponseParameter, err error)
            GetResult(ctx context.Context, serviceID string, messageID string) (result sens.AlimTalkResultResponseParameter, err error)
        }

        var _ AlimtalkClientAPI = (*sens.AlimtalkClient)(nil)
