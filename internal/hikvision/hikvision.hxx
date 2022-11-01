#ifndef Hikvision_H
#define Hikvision_H

#ifdef __cplusplus
extern "C" {
#endif

void test();
unsigned long login(
    char* ip,
    unsigned long port,
    char* username,
    char* password
);
unsigned long getVersion();
unsigned long getBuildVersion();

// typedef void (*MessageCallback)(long lCommand, char* pAlarmInfo, unsigned long dwBufLen, void* pUser);
#ifdef __cplusplus
}
#endif

#endif
