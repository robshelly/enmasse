/*
 * Copyright 2019, EnMasse authors.
 * License: Apache License 2.0 (see the file LICENSE or http://apache.org/licenses/LICENSE-2.0.html).
 */
package io.enmasse.systemtest.framework;

import io.enmasse.systemtest.resolvers.ExtensionContextParameterResolver;
import io.enmasse.systemtest.time.SystemtestsOperation;
import io.enmasse.systemtest.time.TimeMeasuringSystem;
import org.junit.jupiter.api.AfterEach;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.TestInfo;
import org.junit.jupiter.api.extension.ExtendWith;
import org.junit.jupiter.api.extension.ExtensionContext;
import org.slf4j.Logger;

import java.util.Map;

@ExtendWith(ExtensionContextParameterResolver.class)
public interface ITestSeparator {
    Logger testSeparatorLogger = LoggerUtils.getLogger();

    static void printThreadDump() {
        Map<Thread, StackTraceElement[]> allThreads = Thread.getAllStackTraces();
        for (Thread thread : allThreads.keySet()) {
            StringBuilder sb = new StringBuilder();
            StackTraceElement[] trace = allThreads.get(thread);
            sb.append(thread).append("\r\n");
            for (StackTraceElement aTrace : trace) {
                sb.append(" ").append(aTrace).append("\r\n");
            }
            testSeparatorLogger.error(sb.toString());
        }
    }

    @BeforeEach
    default void beforeEachTest(TestInfo testInfo) {
        TimeMeasuringSystem.setTestName(testInfo.getTestClass().get().getName(), testInfo.getDisplayName());
        TimeMeasuringSystem.startOperation(SystemtestsOperation.TEST_EXECUTION);
        LoggerUtils.logDelimiter("#");
        testSeparatorLogger.info("[TEST-START] {}.{}-STARTED", testInfo.getTestClass().get().getName(), testInfo.getDisplayName());
        LoggerUtils.logDelimiter("#");
        testSeparatorLogger.info("");
    }

    @AfterEach
    default void afterEachTest(TestInfo testInfo, ExtensionContext context) {
        if (context.getExecutionException().isPresent()) { // on failed
            Throwable ex = context.getExecutionException().get();
            if (ex instanceof OutOfMemoryError) {
                testSeparatorLogger.error("Got OOM, dumping thread info");
                printThreadDump();
            } else {
                testSeparatorLogger.error("Caught exception", ex);
            }
        }
        TimeMeasuringSystem.stopOperation(SystemtestsOperation.TEST_EXECUTION);
        testSeparatorLogger.info("");
        LoggerUtils.logDelimiter("#");
        testSeparatorLogger.info("[TEST-END] {}.{}-FINISHED", testInfo.getTestClass().get().getName(), testInfo.getDisplayName());
        LoggerUtils.logDelimiter("#");
    }
}