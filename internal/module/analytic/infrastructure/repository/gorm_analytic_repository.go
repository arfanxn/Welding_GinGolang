package repository

import (
	analyticEntity "github.com/arfanxn/welding/internal/module/analytic/domain/entity"
	"github.com/arfanxn/welding/internal/module/analytic/domain/repository"
	mtoEnum "github.com/arfanxn/welding/internal/module/material_test_order/domain/enum"
	roleEnum "github.com/arfanxn/welding/internal/module/role/domain/enum"
	"github.com/arfanxn/welding/internal/module/shared/domain/entity"
	"github.com/arfanxn/welding/pkg/boolutil"
	"github.com/arfanxn/welding/pkg/query"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

type gormAnalyticRepository struct {
	db *gorm.DB
}

type NewGormAnalyticRepositoryParams struct {
	fx.In

	DB *gorm.DB
}

func NewGormAnalyticRepository(params NewGormAnalyticRepositoryParams) repository.AnalyticRepository {
	return &gormAnalyticRepository{db: params.DB}
}

func (r *gormAnalyticRepository) GetSummary(q *query.Query) (*analyticEntity.Summary, error) {
	var (
		userTableName      = entity.NewUser().TableName()
		roleTableName      = entity.NewRole().TableName()
		roleUserTableName  = entity.NewRoleUser().TableName()
		orderTableName     = entity.NewMaterialTestOrder().TableName()
		serviceTableName   = entity.NewMaterialTestService().TableName()
		withDatetimeFilter bool
		summary            analyticEntity.Summary
	)

	args := map[string]any{
		"customer_role_name":             roleEnum.Customer,
		"draft_order_status":             mtoEnum.MaterialTestOrderStatusDraft,
		"awaiting_review_order_status":   mtoEnum.MaterialTestOrderStatusAwaitingReview,
		"rejected_order_status":          mtoEnum.MaterialTestOrderStatusRejected,
		"awaiting_payment_order_status":  mtoEnum.MaterialTestOrderStatusAwaitingPayment,
		"cancelled_order_status":         mtoEnum.MaterialTestOrderStatusCancelled,
		"payment_submitted_order_status": mtoEnum.MaterialTestOrderStatusPaymentSubmitted,
		"payment_rejected_order_status":  mtoEnum.MaterialTestOrderStatusPaymentRejected,
		"payment_approved_order_status":  mtoEnum.MaterialTestOrderStatusPaymentApproved,
		"testing_order_status":           mtoEnum.MaterialTestOrderStatusTesting,
		"refunded_order_status":          mtoEnum.MaterialTestOrderStatusRefunded,
		"completed_order_status":         mtoEnum.MaterialTestOrderStatusCompleted,
	}

	if f := q.GetFilter("date", query.OperatorBetween); f != nil {
		args["start_date"] = f.Values[0]
		args["end_date"] = f.Values[1]
		withDatetimeFilter = true
	}

	err := r.db.Raw(`
		WITH
		-- User counts
		user_counts AS (
			SELECT 
				COUNT(DISTINCT u.id) AS user_count,
				COUNT(DISTINCT u.id) FILTER (WHERE u.activated_at IS NOT NULL AND u.deactivated_at IS NULL) AS active_user_count,
				COUNT(DISTINCT u.id) FILTER (WHERE u.activated_at IS NULL AND u.deactivated_at IS NOT NULL) AS deactive_user_count,
				COUNT(DISTINCT u.id) FILTER (WHERE r.name = @customer_role_name) AS customer_user_count,
				COUNT(DISTINCT u.id) FILTER (WHERE r.name = @customer_role_name AND u.activated_at IS NOT NULL AND u.deactivated_at IS NULL) AS active_customer_user_count,
				COUNT(DISTINCT u.id) FILTER (WHERE r.name = @customer_role_name AND u.activated_at IS NULL AND u.deactivated_at IS NOT NULL) AS deactive_customer_user_count,
				COUNT(DISTINCT u.id) FILTER (WHERE r.name != @customer_role_name) AS employee_user_count,
				COUNT(DISTINCT u.id) FILTER (WHERE r.name != @customer_role_name AND u.activated_at IS NOT NULL AND u.deactivated_at IS NULL) AS active_employee_user_count,
				COUNT(DISTINCT u.id) FILTER (WHERE r.name != @customer_role_name AND u.activated_at IS NULL AND u.deactivated_at IS NOT NULL) AS deactive_employee_user_count
			FROM `+userTableName+` u
			LEFT JOIN `+roleUserTableName+` ru ON u.id = ru.user_id
			LEFT JOIN `+roleTableName+` r ON ru.role_id = r.id
			WHERE 1 = 1`+boolutil.Ternary(withDatetimeFilter, " AND u.created_at BETWEEN @start_date AND @end_date", "")+`
		), 
		-- Order counts
		order_counts AS (
			SELECT 
				COUNT(DISTINCT o.id) AS order_count,
				COUNT(DISTINCT o.id) FILTER (WHERE o.status = @draft_order_status `+boolutil.Ternary(withDatetimeFilter, " AND o.created_at BETWEEN @start_date AND @end_date", "")+`) AS draft_order_count,
				COUNT(DISTINCT o.id) FILTER (WHERE o.status = @awaiting_review_order_status `+boolutil.Ternary(withDatetimeFilter, " AND o.submitted_at BETWEEN @start_date AND @end_date", "")+`) AS awaiting_review_order_count,
				COUNT(DISTINCT o.id) FILTER (WHERE o.status = @rejected_order_status `+boolutil.Ternary(withDatetimeFilter, " AND o.rejected_at BETWEEN @start_date AND @end_date", "")+`) AS rejected_order_count,
				COUNT(DISTINCT o.id) FILTER (WHERE o.status = @awaiting_payment_order_status `+boolutil.Ternary(withDatetimeFilter, " AND o.approved_at BETWEEN @start_date AND @end_date", "")+`) AS awaiting_payment_order_count,
				COUNT(DISTINCT o.id) FILTER (WHERE o.status = @cancelled_order_status `+boolutil.Ternary(withDatetimeFilter, " AND o.cancelled_at BETWEEN @start_date AND @end_date", "")+`) AS cancelled_order_count,
				COUNT(DISTINCT o.id) FILTER (WHERE o.status = @payment_submitted_order_status `+boolutil.Ternary(withDatetimeFilter, " AND o.payment_submitted_at BETWEEN @start_date AND @end_date", "")+`) AS payment_submitted_order_count,
				COUNT(DISTINCT o.id) FILTER (WHERE o.status = @payment_rejected_order_status `+boolutil.Ternary(withDatetimeFilter, " AND o.payment_rejected_at BETWEEN @start_date AND @end_date", "")+`) AS payment_rejected_order_count,
				COUNT(DISTINCT o.id) FILTER (WHERE o.status = @payment_approved_order_status `+boolutil.Ternary(withDatetimeFilter, " AND o.payment_approved_at BETWEEN @start_date AND @end_date", "")+`) AS payment_approved_order_count,
				COUNT(DISTINCT o.id) FILTER (WHERE o.status = @testing_order_status `+boolutil.Ternary(withDatetimeFilter, " AND o.testing_at BETWEEN @start_date AND @end_date", "")+`) AS testing_order_count,
				COUNT(DISTINCT o.id) FILTER (WHERE o.status = @refunded_order_status `+boolutil.Ternary(withDatetimeFilter, " AND o.refunded_at BETWEEN @start_date AND @end_date", "")+`) AS refunded_order_count,
				COUNT(DISTINCT o.id) FILTER (WHERE o.status = @completed_order_status `+boolutil.Ternary(withDatetimeFilter, " AND o.completed_at BETWEEN @start_date AND @end_date", "")+`) AS completed_order_count
			FROM `+orderTableName+` o
		),
		-- Order sum
		order_sum AS (
			SELECT 
				SUM(o.total) AS completed_order_sum
			FROM `+orderTableName+` o
			WHERE o.status = @completed_order_status`+boolutil.Ternary(withDatetimeFilter, " AND o.completed_at BETWEEN @start_date AND @end_date", "")+`
		),
		-- Service counts
		service_count AS (
			SELECT 
				COUNT(DISTINCT s.id) AS service_count
			FROM `+serviceTableName+` s
			WHERE s.deleted_at IS NULL`+boolutil.Ternary(withDatetimeFilter, " AND s.created_at BETWEEN @start_date AND @end_date", "")+`
		)

		SELECT 
			-- User counts
			ucs.user_count,
			ucs.active_user_count,
			ucs.deactive_user_count,
			ucs.customer_user_count,
			ucs.active_customer_user_count,
			ucs.deactive_customer_user_count,
			ucs.employee_user_count,
			ucs.active_employee_user_count,
			ucs.deactive_employee_user_count,
			-- Order counts
			ocs.order_count,
			ocs.draft_order_count,
			ocs.awaiting_review_order_count,
			ocs.rejected_order_count,
			ocs.awaiting_payment_order_count,
			ocs.cancelled_order_count,
			ocs.payment_submitted_order_count,
			ocs.payment_rejected_order_count,
			ocs.payment_approved_order_count,
			ocs.testing_order_count,
			ocs.refunded_order_count,
			ocs.completed_order_count,
			-- Order sum
			os.completed_order_sum,
			-- Service count
			sc.service_count
		FROM user_counts ucs, order_counts ocs, order_sum os, service_count sc
	`, args).Scan(&summary).Error

	if err != nil {
		return nil, err
	}

	return &summary, nil
}

func (r *gormAnalyticRepository) GetOrderTrends(q *query.Query) (analyticEntity.OrderTrends, error) {
	var (
		orderTableName     = entity.NewMaterialTestOrder().TableName()
		withDatetimeFilter bool
		groupByDate        bool
		groupByMonth       bool
		groupByYear        bool
		orderTrends        analyticEntity.OrderTrends
	)

	args := map[string]any{
		"completed_order_status": mtoEnum.MaterialTestOrderStatusCompleted,
	}

	if f := q.GetFilter("date", query.OperatorBetween); f != nil {
		args["start_date"] = f.Values[0]
		args["end_date"] = f.Values[1]
		withDatetimeFilter = true
	}

	if q.GetGroup("year") != nil {
		groupByYear = true
	} else if q.GetGroup("month") != nil {
		groupByMonth = true
	} else {
		groupByDate = true
	}

	sqlString := `
		SELECT 
			` +
		boolutil.Ternary(groupByDate, "DATE(o.completed_at) AS period,", "") +
		boolutil.Ternary(groupByMonth, "TO_CHAR(o.completed_at, 'YYYY-MM') AS period,", "") +
		boolutil.Ternary(groupByYear, "EXTRACT(YEAR FROM o.completed_at)::integer AS period,", "") + `
			COUNT(*) AS count,
			SUM(o.total) AS sum
		FROM ` + orderTableName + ` o
		WHERE o.status = @completed_order_status` +
		boolutil.Ternary(withDatetimeFilter, " AND o.completed_at BETWEEN @start_date AND @end_date", "") + `
		` + boolutil.Ternary(groupByDate, "GROUP BY period ORDER BY period ASC", "") +
		boolutil.Ternary(groupByMonth, "GROUP BY period ORDER BY period ASC", "") +
		boolutil.Ternary(groupByYear, "GROUP BY period ORDER BY period ASC", "")

	err := r.db.Raw(sqlString, args).Scan(&orderTrends).Error

	if err != nil {
		return nil, err
	}

	return orderTrends, nil
}
